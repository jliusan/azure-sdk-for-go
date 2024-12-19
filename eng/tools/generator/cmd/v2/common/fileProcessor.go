// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package common

import (
	"bytes"
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/eng/tools/internal/exports"
	"github.com/Masterminds/semver"
	"golang.org/x/tools/go/ast/astutil"
)

const (
	generated_file_scan_string        = "Code generated by Microsoft"
	autorest_md_file_suffix           = "readme.md"
	autorest_md_module_version_prefix = "module-version: "
	swagger_md_module_name_prefix     = "module-name: "
)

type PullRequestLabel string

const (
	StableLabel                PullRequestLabel = "Stable"
	BetaLabel                  PullRequestLabel = "Beta"
	FirstGALabel               PullRequestLabel = "FirstGA"
	FirstGABreakingChangeLabel PullRequestLabel = "FirstGA,BreakingChange"
	FirstBetaLabel             PullRequestLabel = "FirstBeta"
	StableBreakingChangeLabel  PullRequestLabel = "Stable,BreakingChange"
	BetaBreakingChangeLabel    PullRequestLabel = "Beta,BreakingChange"
)

var (
	v2BeginRegex                    = regexp.MustCompile("^```\\s*yaml\\s*\\$\\(go\\)\\s*&&\\s*\\$\\((track2|v2)\\)")
	v2EndRegex                      = regexp.MustCompile("^\\s*```\\s*$")
	newClientMethodNameRegex        = regexp.MustCompile("^New.*Client$")
	versionLineRegex                = regexp.MustCompile(`moduleVersion\s*=\s*\".*v.+"`)
	changelogPosWithPreviewRegex    = regexp.MustCompile(`##\s*(?P<version>.+)\s*\((\d{4}-\d{2}-\d{2}|Unreleased)\)`)
	changelogPosWithoutPreviewRegex = regexp.MustCompile(`##\s*(?P<version>\d+\.\d+\.\d+)\s*\((\d{4}-\d{2}-\d{2}|Unreleased)\)`)
	packageConfigRegex              = regexp.MustCompile(`\$\((package-.+)\)`)
)

type PackageInfo struct {
	Name        string
	Config      string
	SpecName    string
	RequestLink string
	Tag         string
	ReleaseDate *time.Time
}

// reads from readme.go.md, parses the `track2` section to get module and package name
func ReadV2ModuleNameToGetNamespace(path string) (map[string][]PackageInfo, error) {
	result := make(map[string][]PackageInfo)
	log.Printf("Reading from readme.go.md '%s'...", path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	log.Printf("Parsing module and package name from readme.go.md ...")
	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	var start []int
	var end []int
	for i, line := range lines {
		if v2BeginRegex.MatchString(line) {
			start = append(start, i)
		}
		if len(start) != len(end) && v2EndRegex.MatchString(line) {
			end = append(end, i)
		}
	}

	if len(start) == 0 {
		return nil, fmt.Errorf("cannot find any `track2` section")
	}
	if len(start) != len(end) {
		return nil, fmt.Errorf("last `track2` section does not properly end")
	}

	_, after, _ := strings.Cut(strings.ReplaceAll(path, "\\", "/"), "specification")
	before, _, _ := strings.Cut(after, "resource-manager")
	specName := strings.Trim(before, "/")

	for i := range start {
		hasModuleName := false
		// get the content of the `track2` section
		section := lines[start[i]+1 : end[i]]
		// iterate over the rest lines, get module name
		for _, line := range section {
			if strings.HasPrefix(line, swagger_md_module_name_prefix) {
				modules := strings.Split(strings.TrimSpace(line[len(swagger_md_module_name_prefix):]), "/")
				if len(modules) != 4 {
					return nil, fmt.Errorf("cannot parse module name from `track2` section")
				}
				namespaceName := strings.TrimSuffix(strings.TrimSuffix(modules[3], "\n"), "\r")
				log.Printf("RP: %s Package: %s", modules[2], namespaceName)
				packageConfig := ""
				matchResults := packageConfigRegex.FindAllStringSubmatch(lines[start[i]], -1)
				for _, matchResult := range matchResults {
					packageConfig = matchResult[1] + ": true"
				}
				result[modules[2]] = append(result[modules[2]], PackageInfo{Name: namespaceName, Config: packageConfig, SpecName: specName})
				hasModuleName = true
			}
		}

		if !hasModuleName {
			return nil, fmt.Errorf("%s line:%d-%d is not configured correctly, please refer to sample: `https://github.com/Azure/azure-rest-api-specs/blob/main/documentation/samplefiles/readme.go.md`", path, start[i]+2, end[i])
		}
	}

	return result, nil
}

// remove all sdk generated files in given path
func CleanSDKGeneratedFiles(path string) error {
	log.Printf("Removing all sdk generated files in '%s'...", path)
	return filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(info.Name(), ".go") {
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			if strings.Contains(string(b), generated_file_scan_string) {
				err = os.Remove(path)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}

// replace repo commit with local path in autorest.md file
func ChangeConfigWithLocalPath(path, readmeFile, readmeGoFile string) error {
	log.Printf("Replacing repo commit with local path in autorest.md ...")
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		if strings.Contains(line, autorest_md_file_suffix) {
			lines[i] = fmt.Sprintf("- %s", readmeFile)
			lines[i+1] = fmt.Sprintf("- %s", readmeGoFile)
			break
		}
	}

	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

// replace repo URL and commit id in autorest.md file
func ChangeConfigWithCommitID(path, repoURL, commitID, specRPName string) error {
	log.Printf("Replacing repo URL and commit id in autorest.md ...")
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		if strings.Contains(line, autorest_md_file_suffix) {
			indexResourceManager := strings.Index(line, "resource-manager")
			indexReadme := strings.Index(line, autorest_md_file_suffix)
			resourceManagerPath := []byte(line)
			resourceManagerPath = resourceManagerPath[indexResourceManager : indexReadme-1]

			lines[i] = fmt.Sprintf("- %s/blob/%s/specification/%s/%s/readme.md", repoURL, commitID, specRPName, resourceManagerPath)
			lines[i+1] = fmt.Sprintf("- %s/blob/%s/specification/%s/%s/readme.go.md", repoURL, commitID, specRPName, resourceManagerPath)
			break
		}
	}

	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

// RemoveTagSet delete tag set in config file
func RemoveTagSet(path string) error {
	log.Printf("Removing tag set in autorest.md ...")
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		if strings.Contains(line, "tag:") {
			lines = append(lines[:i], lines[i+1:]...)
			break
		}
	}

	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

// get swagger rp folder name from autorest.md file
func GetSpecRpName(packageRootPath string) (string, error) {
	b, err := os.ReadFile(path.Join(packageRootPath, "autorest.md"))
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if strings.Contains(line, autorest_md_file_suffix) {
			allParts := strings.Split(line, "/")
			for i, part := range allParts {
				if part == "specification" {
					return allParts[i+1], nil
				}
			}
		}
	}
	return "", fmt.Errorf("cannot get sepc rp name from config")
}

// replace version: use `module-version: ` prefix to locate version in autorest.md file, use version = "v*.*.*" regrex to locate version in constants.go file
func ReplaceVersion(packageRootPath string, newVersion string) error {
	path := filepath.Join(packageRootPath, "autorest.md")
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, autorest_md_module_version_prefix) {
			lines[i] = line[:len(autorest_md_module_version_prefix)] + newVersion
			break
		}
	}

	if err = os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644); err != nil {
		return err
	}

	path = filepath.Join(packageRootPath, "constants.go")
	if b, err = os.ReadFile(path); err != nil {
		return err
	}
	contents := versionLineRegex.ReplaceAllString(string(b), "moduleVersion = \"v"+newVersion+"\"")

	return os.WriteFile(path, []byte(contents), 0644)
}

// calculate new version by changelog using semver package
func CalculateNewVersion(changelog *Changelog, previousVersion string, isCurrentPreview bool) (*semver.Version, PullRequestLabel, error) {
	version, err := semver.NewVersion(previousVersion)
	if err != nil {
		return nil, "", err
	}
	log.Printf("Lastest version is: %s", version.String())

	var newVersion semver.Version
	var prl PullRequestLabel
	if version.Major() == 0 {
		// preview version calculation
		if !isCurrentPreview {
			tempVersion, err := semver.NewVersion("1.0.0")
			if err != nil {
				return nil, "", err
			}
			newVersion = *tempVersion
			if changelog.HasBreakingChanges() {
				prl = FirstGABreakingChangeLabel
			} else {
				prl = FirstGALabel
			}
		} else if changelog.HasBreakingChanges() {
			newVersion = version.IncMinor()
			prl = BetaBreakingChangeLabel
		} else if changelog.Modified.HasAdditiveChanges() {
			newVersion = version.IncMinor()
			prl = BetaLabel
		} else {
			newVersion = version.IncPatch()
			prl = BetaLabel
		}
	} else {
		if isCurrentPreview {
			if strings.Contains(previousVersion, "beta") {
				betaNumber, err := strconv.Atoi(strings.Split(version.Prerelease(), "beta.")[1])
				if err != nil {
					return nil, "", err
				}
				newVersion, err = version.SetPrerelease("beta." + strconv.Itoa(betaNumber+1))
				if err != nil {
					return nil, "", err
				}
				if changelog.HasBreakingChanges() {
					prl = BetaBreakingChangeLabel
				} else {
					prl = BetaLabel
				}
			} else {
				if changelog.HasBreakingChanges() {
					newVersion = version.IncMajor()
					prl = BetaBreakingChangeLabel
				} else if changelog.Modified.HasAdditiveChanges() {
					newVersion = version.IncMinor()
					prl = BetaLabel
				} else {
					newVersion = version.IncPatch()
					prl = BetaLabel
				}
				newVersion, err = newVersion.SetPrerelease("beta.1")
				if err != nil {
					return nil, "", err
				}
			}
		} else {
			if strings.Contains(previousVersion, "beta") {
				return nil, "", fmt.Errorf("must have stable previous version")
			}
			// release version calculation
			if changelog.HasBreakingChanges() {
				newVersion = version.IncMajor()
				prl = StableBreakingChangeLabel
			} else if changelog.Modified.HasAdditiveChanges() {
				newVersion = version.IncMinor()
				prl = StableLabel
			} else {
				newVersion = version.IncPatch()
				prl = StableLabel
			}
		}
	}

	log.Printf("New version is: %s", newVersion.String())
	return &newVersion, prl, nil
}

// add new changelog md to changelog file
func AddChangelogToFile(changelog *Changelog, version *semver.Version, packageRootPath, releaseDate string) (string, error) {
	path := filepath.Join(packageRootPath, ChangelogFileName)
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	oldChangelog := string(b)
	newChangelog := "# Release History\n\n"
	matchResults := changelogPosWithPreviewRegex.FindAllStringSubmatchIndex(oldChangelog, -1)
	additionalChangelog := changelog.ToCompactMarkdown()
	if releaseDate == "" {
		releaseDate = time.Now().Format("2006-01-02")
	}

	for _, matchResult := range matchResults {
		newChangelog = newChangelog + "## " + version.String() + " (" + releaseDate + ")\r\n" + additionalChangelog + "\r\n\r\n" + oldChangelog[matchResult[0]:]
		break
	}

	err = os.WriteFile(path, []byte(newChangelog), 0644)
	if err != nil {
		return "", err
	}
	return additionalChangelog, nil
}

// replace `{{NewClientName}}` placeholder in README.md by first func name according to `^New.+Method$` pattern
func ReplaceNewClientNamePlaceholder(packageRootPath string, exports exports.Content) error {
	path := filepath.Join(packageRootPath, "README.md")
	var clientName string
	for _, k := range SortFuncItem(exports.Funcs) {
		v := exports.Funcs[k]
		if newClientMethodNameRegex.MatchString(k) && *v.Params == "string, azcore.TokenCredential, *arm.ClientOptions" {
			clientName = k
			break
		}
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cannot read from file '%s': %+v", path, err)
	}

	var content = strings.ReplaceAll(string(b), "{{NewClientName}}", clientName)
	return os.WriteFile(path, []byte(content), 0644)
}

func UpdateModuleDefinition(packageRootPath, relativePath string, version *semver.Version) error {
	if version.Major() > 1 {
		path := filepath.Join(packageRootPath, "go.mod")

		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("cannot parse version from changelog")
		}

		lines := strings.Split(string(b), "\n")
		for i, line := range lines {
			if strings.HasPrefix(line, "module") {
				line = strings.TrimRight(line, "\r")
				parts := strings.Split(line, "/")
				if parts[len(parts)-1] != fmt.Sprintf("v%d", version.Major()) {
					lines[i] = fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/%s/v%d", relativePath, version.Major())
				}
				break
			}
		}
		if err = os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644); err != nil {
			return err
		}
	}
	return nil
}

func UpdateOnboardChangelogVersion(packageRootPath, versionNumber string) error {
	changelogPath := filepath.Join(packageRootPath, ChangelogFileName)
	b, err := os.ReadFile(changelogPath)
	if err != nil {
		return err
	}

	newChangelog := regexp.MustCompile(`\d\.\d\.\d`).ReplaceAllString(string(b), versionNumber)
	err = os.WriteFile(changelogPath, []byte(newChangelog), 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetAlwaysSetBodyParamRequiredFlag(path string) (string, error) {
	buildGo, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	if strings.Contains(string(buildGo), "-alwaysSetBodyParamRequired") {
		return "-alwaysSetBodyParamRequired", nil
	}
	return "", nil
}

// AddTagSet add tag in file
func AddTagSet(path, tag string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		if strings.Contains(line, "tag:") {
			lines[i] = tag
			break
		}

		// end index
		if i == len(lines)-1 {
			for j := len(lines) - 1; j > 0; j-- {
				if strings.Contains(lines[j], "```") {
					if lines[j-1] == "" {
						lines[j-1] = tag
						break
					} else {
						newLines := make([]string, len(lines))
						copy(newLines, lines)

						newLines = append(newLines[:j], tag)
						tailLines := lines[j:]
						lines = append(newLines, tailLines...)
						break
					}
				}
			}
		}
	}

	return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

func GetTag(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if strings.Contains(line, "tag:") {
			return strings.TrimSpace(string([]byte(line)[len("tag:"):])), nil
		}
	}

	return "", nil
}

func replaceModuleImport(path, rpName, namespaceName, previousVersion, currentVersion, subPath string, suffixes ...string) error {
	previous, err := semver.NewVersion(previousVersion)
	if err != nil {
		return err
	}

	current, err := semver.NewVersion(currentVersion)
	if err != nil {
		return err
	}

	if previous.Major() == current.Major() {
		return nil
	}

	oldModule := fmt.Sprintf("github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/%s/%s", rpName, namespaceName)
	if previous.Major() > 1 {
		oldModule = fmt.Sprintf("%s/v%d", oldModule, previous.Major())
	}

	newModule := fmt.Sprintf("github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/%s/%s", rpName, namespaceName)
	if current.Major() > 1 {
		newModule = fmt.Sprintf("%s/v%d", newModule, current.Major())
	}

	if oldModule == newModule {
		return nil
	}

	return filepath.Walk(filepath.Join(path, subPath), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		suffix := false
		for i := 0; i < len(suffixes) && !suffix; i++ {
			suffix = strings.HasSuffix(info.Name(), suffixes[i])
		}

		if suffix {
			b, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			newFile := strings.ReplaceAll(string(b), fmt.Sprintf("\"%s\"", oldModule), fmt.Sprintf("\"%s\"", newModule))
			if newFile != string(b) {
				if err = os.WriteFile(path, []byte(newFile), 0666); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func getModuleVersion(autorestPath string) (*semver.Version, error) {
	data, err := os.ReadFile(autorestPath)
	if err != nil {
		return nil, err
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, autorest_md_module_version_prefix) {
			return semver.NewVersion(strings.TrimSpace(line[len(autorest_md_module_version_prefix):]))
		}
	}

	return nil, errors.New("module-version does not exist in autorest.md")
}

func existSuffixFile(path, suffix string) bool {

	existed := false
	err := filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(d.Name(), suffix) {
			existed = true
		}
		return nil
	})
	if err != nil {
		return false
	}

	return existed
}

func replaceReadmeModule(path, relativePath, currentVersion string) error {
	readmeFile, err := os.ReadFile(filepath.Join(path, "README.md"))
	if err != nil {
		return err
	}

	module := fmt.Sprintf("github.com/Azure/azure-sdk-for-go/%s", relativePath)

	readmeModule := module
	match := regexp.MustCompile(fmt.Sprintf(`%s/v\d+`, module))
	if match.Match(readmeFile) {
		readmeModule = match.FindString(string(readmeFile))
	}

	newModule := module
	current, err := semver.NewVersion(currentVersion)
	if err != nil {
		return err
	}
	if current.Major() > 1 {
		newModule = fmt.Sprintf("%s/v%d", newModule, current.Major())
	}

	if newModule == readmeModule {
		return nil
	}

	newReadmeFile := bytes.ReplaceAll(readmeFile, []byte(readmeModule), []byte(newModule))
	err = os.WriteFile(filepath.Join(path, "README.md"), newReadmeFile, 0666)
	if err != nil {
		return err
	}

	return nil
}

func ReplaceReadmeNewClientName(packageRootPath string, exports exports.Content) error {
	path := filepath.Join(packageRootPath, "README.md")
	var clientName string
	for _, k := range SortFuncItem(exports.Funcs) {
		v := exports.Funcs[k]
		if newClientMethodNameRegex.MatchString(k) && *v.Params == "string, azcore.TokenCredential, *arm.ClientOptions" {
			clientName = k
			break
		}
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cannot read from file '%s': %+v", path, err)
	}

	oldClientName := ""
	for _, v := range strings.Split(string(b), "\n") {
		oldClientName = regexp.MustCompile(`New.*Client\(\)`).FindString(v)
		if oldClientName != "" {
			break
		}
	}

	if clientName == "" || oldClientName == "" || oldClientName == fmt.Sprintf("%s()", clientName) {
		return nil
	}

	var content = strings.ReplaceAll(string(b), oldClientName, fmt.Sprintf("%s()", clientName))
	return os.WriteFile(path, []byte(content), 0644)
}

func ReplaceConstModuleVersion(packagePath string, newVersion string) error {
	path := filepath.Join(packagePath, "constants.go")
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	contents := versionLineRegex.ReplaceAllString(string(data), "moduleVersion = \"v"+newVersion+"\"")
	if contents == string(data) {
		return nil
	}

	return os.WriteFile(path, []byte(contents), 0644)
}

func ReplaceModule(newVersion *semver.Version, packagePath, baseModule string, suffixs ...string) error {
	return filepath.Walk(packagePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		hasSuffix := slices.ContainsFunc(suffixs, func(s string) bool { return strings.HasSuffix(info.Name(), s) })
		if len(suffixs) == 0 || hasSuffix {
			if err = ReplaceImport(path, baseModule, newVersion.Major()); err != nil {
				return err
			}
		}

		return nil
	})
}

func ReplaceImport(sourceFile string, baseModule string, majorVersion int64) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, sourceFile, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	rewrote := false
	for _, i := range f.Imports {
		if strings.HasPrefix(i.Path.Value, fmt.Sprintf("\"%s", baseModule)) {
			oldPath := importPath(i)
			after, _ := strings.CutPrefix(oldPath, baseModule)

			newPath := baseModule
			if after != "" {
				before, sub, _ := strings.Cut(strings.TrimLeft(after, "/"), "/")
				if regexp.MustCompile(`^v\d+$`).MatchString(before) {
					if majorVersion > 1 {
						newPath = fmt.Sprintf("%s/v%d", baseModule, majorVersion)
					}
					if sub != "" {
						newPath = fmt.Sprintf("%s/%s", newPath, sub)
					}
				} else {
					if majorVersion > 1 {
						newPath = fmt.Sprintf("%s/v%d", baseModule, majorVersion)
					}
					newPath = fmt.Sprintf("%s/%s", newPath, before)
					if sub != "" {
						newPath = fmt.Sprintf("%s/%s", newPath, sub)
					}
				}
			} else {
				if majorVersion > 1 {
					newPath = fmt.Sprintf("%s/v%d", baseModule, majorVersion)
				}
			}

			if newPath != oldPath {
				rewrote = astutil.RewriteImport(fset, f, oldPath, newPath)
			}
		}
	}

	if rewrote {
		w, err := os.Create(sourceFile)
		if err != nil {
			return err
		}
		defer w.Close()

		return printer.Fprint(w, fset, f)
	}

	return nil
}

func importPath(s *ast.ImportSpec) string {
	t, err := strconv.Unquote(s.Path.Value)
	if err != nil {
		return ""
	}
	return t
}
