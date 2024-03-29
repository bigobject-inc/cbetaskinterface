# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).


## [v1.6.1] - 2022-07-06

### Changed

- use go 1.18.
- instance ID become string.


## [v1.6.0] - 2022-05-24

### Added

- add LicenseSetting


## [v1.5.6] - 2022-03-01

### Changed

- use stailib@v1.8.1

## [v1.5.4] - 2022-01-20

### Changed

- use stailib@v1.5.4


## [v1.5.3] - 2021-09-15

### Changed

- github.com/bigobject-inc/stailib v1.6.0
- use stailib.cs.APISetting.
- 

## [v1.5.1] - 2021-07-30

### Changed

- remove []ttType.Node, use EventData interface{}.
- remove RaiseExternalFunc.


## [v1.4.2] - 2021-04-07

### Add

- add Mode in CBEEngineSetting.


## [v1.4.1] - 2021-03-25

### Changed

- remove cbe-task-interface.go and cbe-task-type.go
- add CBEEngine and CBEEngineSetting.
- remove TaskSetting and Task.
- Process(cbe.CBE, []*Event) ([]*Event, error)

## [v1.3.1] - 2021-03-09

### Changed

- add type TaskSetting
- add Defer()
- change Init()
- add LogLevel

## [v1.2.2] - 2021-02-18

### Changed

- use github.com/bigobject-inc/stailib v1.4.0
- use cbe-utility v3.2.0
- Event.toSequenceName -> Event.toInstanceName

## [v1.1.3] - 2021-02-17

### Add

- add Event struct
- change Prepare(CBE) -> Prepare(CBE, []*Event)
- change Process() (error) -> Process() ([]*Event)


## [v1.0.10] - 2020-12-29

### Add

- add GetVersion

## [v1.0.9] - 2020-12-22

### Change

- change ownership to bigobject-inc
- change go.mod


## [v1.0.3] - 2020-12-22

### Change

- simplified, only left Task interface

## [v1.0.2] - 2020-12-22

### Change

- change go.mod

## [v1.0.1] - 2020-12-21

### Change

- export DefineSequence and DefineSet


## [v1.0.0] - 2020-12-21

### Add

- Add cbetaskinterface


[v1.6.1]: github.com/bigobject-inc/cbetaskinterface/archive/v1.6.1
[v1.6.0]: github.com/bigobject-inc/cbetaskinterface/archive/v1.6.0
[v1.5.6]: github.com/bigobject-inc/cbetaskinterface/archive/v1.5.6
[v1.5.4]: github.com/bigobject-inc/cbetaskinterface/archive/v1.5.4
[v1.5.3]: github.com/bigobject-inc/cbetaskinterface/archive/v1.5.3
[v1.5.1]: github.com/bigobject-inc/cbetaskinterface/archive/v1.5.1
[v1.4.2]: github.com/bigobject-inc/cbetaskinterface/archive/v1.4.2
[v1.4.1]: github.com/bigobject-inc/cbetaskinterface/archive/v1.4.1
[v1.3.0]: github.com/bigobject-inc/cbetaskinterface/archive/v1.3.0
[v1.2.2]: github.com/bigobject-inc/cbetaskinterface/archive/v1.2.2
[v1.1.3]: github.com/bigobject-inc/cbetaskinterface/archive/v1.1.3
[v1.0.10]: github.com/bigobject-inc/cbetaskinterface/archive/v1.0.10
[v1.0.9]: github.com/bigobject-inc/cbetaskinterface/archive/v1.0.9
[v1.0.1]: github.com/bigobject-inc/cbetaskinterface/archive/v1.0.1
[v1.0.0]: github.com/bigobject-inc/cbetaskinterface/archive/v1.0.0

