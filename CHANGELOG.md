# Change Log
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/) 
and this project adheres to [Semantic Versioning](http://semver.org/).
## [Unreleased]
### Added
### Changed
## [0.4.1]
### Added
### Changed
* library updates
* switching to ginkgo v2
## [0.4.0]
### Added
### Changed
Allow nun numeric buildcounters. They will be added as strings to the prerelease versions. 
This allows git short hashes in docker labels
### Fixed
All release candidates now use the -rc denotion in their semver
## [0.3.1]
### Added
### Changed
### Fixed
All release candidates now use the -rc denotion in their semver
## [0.3.0] 2018-12-21
### Added
Export a version which can be used in go module setups. (includes the no standard v prefix)
Use ```goModuleBuildVersion``` to tag go module repositories 
### Changed
## [0.2.2] 2018-12-20
### Added
Switch to use go modules
### Changed
## [0.2.1] 2017-01-6
### Added
Buildhelper: export a variable with s stripped metadata so it can be used as a kubernetes label (Kubernetes does not accept complete semvars)
## [0.2.0] 2017-01-5
### Added
- manipulator: allows to bump versions and set prerelease and build metadata
- projectversion service which allows init of a repository
    + next 
    + release

- cli : init call
- cli : next [patch|minor|major] [prerelease tags]
- cli : release makes a release
- cli : generate; writes a helper file for ci 

### Changed
- cli is now located in the cmd folder
- version file is now called version.yml

## [0.1.0] 2016-12-30
### Added
- show command: shows the current semver version

### Changed

