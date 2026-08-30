# Companies dataset

`companies_dataset.json` is a manually-vendored snapshot mapping LeetCode
question numbers to companies known to have asked them (as JSON object
keys are strings, question numbers are quoted: `"965": ["Google", "Amazon"]`).

LeetCode's own company-tag data is Premium-only and has no public API, so
this file starts empty and is refreshed by hand — periodically replace it
wholesale with a fresh export from a public community dataset (e.g. a
GitHub repo tracking company-wise LeetCode question lists). The
content-gen agent's `COMPANIES.md` output is best-effort: if a question
number has no entry here, that file is simply not generated.
