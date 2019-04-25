/* Copyright 2019 The Bazel Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gazelle

import (
	"errors"
	"fmt"
	"log"
	"path"
	"sort"
	"strings"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/repo"
	"github.com/bazelbuild/bazel-gazelle/resolve"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

var _ = fmt.Printf

var (
	skipImportError = errors.New("std import")
	notFoundError   = errors.New("not found")
)

// Name returns the name of the language. This should be a prefix of the
// kinds of rules generated by the language, e.g., "go" for the Go extension
// since it generates "go_library" rules.
func (s *sasslang) Name() string {
	return "js"
}

// Imports returns a list of ImportSpecs that can be used to import the rule
// r. This is used to populate RuleIndex.
//
// If nil is returned, the rule will not be indexed. If any non-nil slice is
// returned, including an empty slice, the rule will be indexed.
func (s *sasslang) Imports(c *config.Config, r *rule.Rule, f *rule.File) []resolve.ImportSpec {
	rel := f.Pkg
	srcs := r.AttrStrings("srcs")
	imports := make([]resolve.ImportSpec, len(srcs))
	for i, src := range srcs {
		withoutSuffix := strings.TrimSuffix(src, path.Ext(src))
		imports[i] = resolve.ImportSpec{
			Lang: "js",
			Imp:  strings.ToLower(path.Join("@/", rel, withoutSuffix)),
		}
	}
	return imports
}

// Embeds returns a list of labels of rules that the given rule embeds. If
// a rule is embedded by another importable rule of the same language, only
// the embedding rule will be indexed. The embedding rule will inherit
// the imports of the embedded rule.
func (s *sasslang) Embeds(r *rule.Rule, from label.Label) []label.Label {
	// Sass doesn't have a concept of embedding as far as I know.
	return nil
}

// Resolve translates imported libraries for a given rule into Bazel
// dependencies. A list of imported libraries is typically stored in a
// private attribute of the rule when it's generated (this interface doesn't
// dictate how that is stored or represented). Resolve generates a "deps"
// attribute (or the appropriate language-specific equivalent) for each
// import according to language-specific rules and heuristics.
func (s *sasslang) Resolve(c *config.Config, ix *resolve.RuleIndex, rc *repo.RemoteCache, r *rule.Rule, importsRaw interface{}, from label.Label) {
	imports := importsRaw.([]string)
	r.DelAttr("deps")
	depSet := make(map[string]bool)
	for _, imp := range imports {
		log.Printf("imp: %v, ix: %v, from %v", imp, ix, from)
		l, err := resolveWithIndex(ix, imp, from)
		if err == skipImportError {
			continue
		} else if err == notFoundError {
			log.Printf("from.Repo: %v, from.Pkg: %v, l: %v", from.Repo, from.Pkg, l.Rel(from.Repo, from.Pkg).String())
			log.Printf("Import not found: Imports: %v\n", imp)
		} else if err != nil {
			log.Print(err)
		} else {
			l = l.Rel(from.Repo, from.Pkg)
			log.Printf("from.Repo: %v, from.Pkg: %v, l: %v", from.Repo, from.Pkg, l.Rel(from.Repo, from.Pkg).String())
			depSet[l.String()] = true
		}
	}
	if len(depSet) > 0 {
		deps := make([]string, 0, len(depSet))
		for dep := range depSet {
			deps = append(deps, dep)
		}
		sort.Strings(deps)
		r.SetAttr("deps", deps)
	}
}

func resolveWithIndex(ix *resolve.RuleIndex, imp string, from label.Label) (label.Label, error) {
	res := resolve.ImportSpec{
		Lang: "js",
		Imp:  imp,
	}
	log.Printf("my res %v", res)
	matches := ix.FindRulesByImport(res, "js")
	log.Printf("matches: %v", matches)
	if len(matches) == 0 {
		return label.NoLabel, notFoundError
	}
	if len(matches) > 1 {
		return label.NoLabel, fmt.Errorf("multiple rules (%s and %s) may be imported with %q from %s", matches[0].Label, matches[1].Label, imp, from)
	}
	if matches[0].IsSelfImport(from) {
		return label.NoLabel, skipImportError
	}
	return matches[0].Label, nil
}
