# SPS Buddy
A tool to check coding standards in structured text.
# Usage
```
sps-buddy example.scl
```

Add steps like these to your github actions.
```
- name: Install SPS-Buddy
  uses: jaxxstorm/action-install-gh-release@v1.10.0 
  with:
    repo: hhirsch/sps-buddy
    chmod: 0755
- name: Check Code-Conventions
  run: sps-buddy --batch > /dev/null
```
# Features
- checks if your variable names are in mixed camel case
- returns proper exit codes for use in CI pipelines
- output is routet to stdout and stderr so you can handle error messages and regular output separately

SPS buddy is free Software licensed under the GNU General Public License v3. 
See <http://www.gnu.org/licenses/gpl-3.0.html> for details.
Copyright (C) 2025  Henry Hirsch
