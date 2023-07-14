# deptree

show golang dependence like tree

## How to get it

```shell
➜ go install github.com/vc60er/deptree@master
➜ export PATH=$PATH:$GOPATH/bin
```

## Usage

### params

```text
Usage of deptree:
  -a  show all dependencies, also without upgrade and point out duplicated children
  -c  upgrade candidates will be marked yellow
  -d int
      max depth of dependencies (default 3)
  -f  force show of each occurrence of a child branch in tree (can cause hang)
  -graph string
      path to file created e.g. by 'go mod graph > grapphfile.txt'
  -json
      print JSON instead of tree
  -t  visualize trimmed tree by '└─...'
  -upgrade string
      path to file created e.g. by 'go list -u -m -json all > upgradefile.txt'
  -v int
      be more verbose
```

### example

```shell
➜  redis git:(master) go mod graph | deptree -d 2 -t -a
call 'go list -u -m -json all', be patient...
dependency tree with depth 2 for package: gobot.io/x/gobot, least 99 trimmed item(s)
* tree of duplicate children not drawn (-f not set)
* 12 duplicate child trees replaced by links

gobot.io/x/gobot (go1.17)
 ├── github.com/bmizerany/pat@v0.0.0-20210406213842-e4b6760bdd6f
 ├── github.com/cpuguy83/go-md2man/v2@v2.0.2 (go1.11)
 │    └── github.com/russross/blackfriday/v2@v2.1.0
 ├── github.com/creack/goselect@v0.1.2 (go1.13)
 ├── github.com/davecgh/go-spew@v1.1.1
 ├── github.com/donovanhide/eventsource@v0.0.0-20210830082556-c59027999da0
 ├── github.com/eclipse/paho.mqtt.golang@v1.4.1 (go1.14) => [v1.4.2]
 │    ├── github.com/gorilla/websocket@v1.4.2 => [v1.5.0] (go1.12)
 │    ├── golang.org/x/net@v0.0.0-20200425230154-ff2c4b7c35a0 => [v0.10.0]
 │    │    └── 3 more ...
 │    └── golang.org/x/sync@v0.0.0-20210220032951-036812b2e83c => [v0.2.0]
 ├── github.com/fatih/structs@v1.1.0
 ├── <2> github.com/go-ole/go-ole@v1.2.6 (go1.12)
 │    └── golang.org/x/sys@v0.0.0-20190916202348-b4ddaad3f8a3 => [v0.8.0]
 ├── github.com/godbus/dbus/v5@v5.1.0 (go1.12)
 ├── github.com/gofrs/uuid@v4.3.0+incompatible => [v4.4.0+incompatible]
 ├── <7> github.com/golang/protobuf@v1.5.0 (go1.9) => [v1.5.3]
 │    ├── <8> github.com/google/go-cmp@v0.5.5 (go1.8) => [v0.5.9]
 │    │    └── 1 more ...
 │    └── google.golang.org/protobuf@v1.26.0-rc.1 => [v1.30.0]
 │         └── 1 more ...
 ├── github.com/gorilla/websocket@v1.5.0 (go1.12)
 ├── github.com/hashicorp/errwrap@v1.1.0
 ├── github.com/hashicorp/go-multierror@v1.1.1 (go1.13)
 │    └── github.com/hashicorp/errwrap@v1.0.0 => [v1.1.0]
 ├── github.com/hybridgroup/go-ardrone@v0.0.0-20140402002621-b9750d8d7b78
 ├── github.com/hybridgroup/mjpeg@v0.0.0-20140228234708-4680f319790e
 ├── <10> github.com/muka/go-bluetooth@v0.0.0-20220830075246-0746e3a1ea53 (go1.14) => [v0.0.0-20221213043340-85dc80edc4e1]
 │    ├── github.com/fatih/structs@v1.1.0
 │    ├── github.com/godbus/dbus/v5@v5.0.3 => [v5.1.0] (go1.12)
 │    ├── github.com/niemeyer/pretty@v0.0.0-20200227124842-a10e7caefd8e (go1.12)
 │    │    └── 1 more ...
 │    ├── github.com/pkg/errors@v0.9.1
 │    ├── github.com/sirupsen/logrus@v1.6.0 => [v1.9.0] (go1.13)
 │    │    └── 5 more ...
 │    ├── github.com/stretchr/testify@v1.6.1 => [v1.8.2]
 │    │    └── 4 more ...
 │    ├── github.com/suapapa/go_eddystone@v1.3.1 (go1.14)
 │    │    └── 2 more ...
 │    ├── golang.org/x/sys@v0.0.0-20200728102440-3e129f6d46b1 => [v0.8.0]
 │    ├── golang.org/x/tools@v0.0.0-20200925191224-5d1fdd8fa346 => [v0.9.1]
 │    │    └── 5 more ...
 │    └── gopkg.in/check.v1@v1.0.0-20200227125254-8fa46927fb4f => [v1.0.0-20201130134442-10cb98267c6c]
 ├── github.com/nats-io/nats-server/v2@v2.1.0 => [v2.9.16]
 │    ├── github.com/golang/protobuf@v1.3.2 => [v1.5.3]
 │    ├── github.com/nats-io/jwt@v0.3.0 => [v1.2.2]
 │    │    └── 1 more ...
 │    ├── github.com/nats-io/nats.go@v1.8.1 => [v1.25.0]
 │    │    └── 2 more ...
 │    ├── github.com/nats-io/nkeys@v0.1.0 => [v0.4.4]
 │    │    └── 1 more ...
 │    ├── github.com/nats-io/nuid@v1.0.1
 │    ├── golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4 => [v0.9.0]
 │    │    └── 2 more ...
 │    └── golang.org/x/sys@v0.0.0-20190726091711-fc99dfbffb4e => [v0.8.0]
 ├── github.com/nats-io/nats.go@v1.18.0 (go1.16) => [v1.25.0]
 │    ├── github.com/nats-io/nkeys@v0.3.0 (go1.16) => [v0.4.4], see <1> for 1 children (branch 'gobot.io/x/gobot' level 1)
 │    └── github.com/nats-io/nuid@v1.0.1
 ├── <1> github.com/nats-io/nkeys@v0.3.0 (go1.16) => [v0.4.4]
 │    └── golang.org/x/crypto@v0.0.0-20210314154223-e6e6c4f2bb5b => [v0.9.0]
 │         └── 3 more ...
 ├── github.com/nats-io/nuid@v1.0.1
 ├── github.com/pkg/errors@v0.9.1
 ├── github.com/pmezard/go-difflib@v1.0.0
 ├── github.com/russross/blackfriday/v2@v2.1.0
 ├── github.com/saltosystems/winrt-go@v0.0.0-20220913104103-712830fcd2ad (go1.18) => [v0.0.0-20230510070731-e096b9afa761]
 │    ├── github.com/glerchundi/subcommands@v0.0.0-20181212083838-923a6ccb11f8
 │    ├── github.com/go-kit/log@v0.2.1 (go1.17)
 │    │    └── 1 more ...
 │    ├── github.com/go-ole/go-ole@v1.2.6 (go1.12), see <2> for 1 children (branch 'gobot.io/x/gobot' level 1)
 │    ├── github.com/peterbourgon/ff/v3@v3.1.2 (go1.16) => [v3.3.1]
 │    │    └── 2 more ...
 │    ├── github.com/stretchr/testify@v1.7.5 => [v1.8.2]
 │    │    └── 4 more ...
 │    ├── github.com/tdakkota/win32metadata@v0.1.0 (go1.16)
 │    │    └── 2 more ...
 │    ├── golang.org/x/tools@v0.1.11 (go1.17) => [v0.9.1]
 │    │    └── 6 more ...
 │    ├── github.com/davecgh/go-spew@v1.1.1
 │    ├── github.com/go-logfmt/logfmt@v0.5.1 (go1.17) => [v0.6.0]
 │    ├── github.com/pmezard/go-difflib@v1.0.0
 │    ├── golang.org/x/mod@v0.6.0-dev.0.20220419223038-86c51ed26bb4 (go1.17) => [v0.10.0]
 │    │    └── 2 more ...
 │    ├── golang.org/x/sys@v0.0.0-20220624220833-87e55d714810 => [v0.8.0]
 │    └── gopkg.in/yaml.v3@v3.0.1, see <3> for 1 children (branch 'gobot.io/x/gobot' level 1)
 ├── github.com/sigurn/crc8@v0.0.0-20220107193325-2243fe600f9f (go1.17)
 ├── <11> github.com/sirupsen/logrus@v1.9.0 (go1.13)
 │    ├── github.com/davecgh/go-spew@v1.1.1
 │    ├── <5> github.com/stretchr/testify@v1.7.0 => [v1.8.2]
 │    │    └── 4 more ...
 │    └── golang.org/x/sys@v0.0.0-20220715151400-c0bba94af5f8 => [v0.8.0]
 ├── github.com/stretchr/testify@v1.8.0 (go1.13) => [v1.8.2]
 │    ├── github.com/davecgh/go-spew@v1.1.1
 │    ├── github.com/pmezard/go-difflib@v1.0.0
 │    ├── github.com/stretchr/objx@v0.4.0 (go1.12) => [v0.5.0]
 │    │    └── 2 more ...
 │    └── gopkg.in/yaml.v3@v3.0.1, see <3> for 1 children (branch 'gobot.io/x/gobot' level 1)
 ├── <12> github.com/tinygo-org/cbgo@v0.0.4 (go1.14)
 │    └── github.com/sirupsen/logrus@v1.5.0 => [v1.9.0] (go1.13)
 │         └── 5 more ...
 ├── github.com/urfave/cli@v1.22.10 (go1.11) => [v1.22.13]
 │    ├── github.com/BurntSushi/toml@v0.3.1 => [v1.2.1]
 │    ├── github.com/cpuguy83/go-md2man/v2@v2.0.0-20190314233015-f79a8a8ca69d => [v2.0.2] (go1.11)
 │    │    └── 3 more ...
 │    └── <4> gopkg.in/yaml.v2@v2.2.2 => [v2.4.0]
 │         └── 1 more ...
 ├── github.com/veandco/go-sdl2@v0.4.25 (go1.15) => [v0.4.35]
 ├── github.com/warthog618/gpiod@v0.8.0 (go1.17) => [v0.8.1]
 │    ├── github.com/pilebones/go-udev@v0.0.0-20180820235104-043677e09b13 => [v0.9.0]
 │    ├── github.com/spf13/cobra@v0.0.5 (go1.12) => [v1.7.0]
 │    ├── github.com/spf13/pflag@v1.0.5 (go1.12)
 │    ├── github.com/stretchr/testify@v1.4.0 => [v1.8.2]
 │    ├── github.com/warthog618/config@v0.4.1 (go1.11) => [v0.5.1]
 │    ├── golang.org/x/sys@v0.0.0-20200223170610-d5e6a3e2c0ae => [v0.8.0]
 │    ├── gopkg.in/check.v1@v1.0.0-20190902080502-41f04d3bba15 => [v1.0.0-20201130134442-10cb98267c6c]
 │    ├── github.com/davecgh/go-spew@v1.1.1
 │    ├── github.com/fsnotify/fsnotify@v1.4.7 => [v1.6.0]
 │    ├── github.com/inconshreveable/mousetrap@v1.0.0 => [v1.1.0]
 │    ├── github.com/pkg/errors@v0.8.1 => [v0.9.1]
 │    ├── github.com/pmezard/go-difflib@v1.0.0
 │    └── gopkg.in/yaml.v2@v2.2.2 => [v2.4.0], see <4> for 1 children (branch 'github.com/urfave/cli@v1.22.10' level 2)
 ├── go.bug.st/serial@v1.4.0 (go1.17) => [v1.5.0]
 │    ├── github.com/creack/goselect@v0.1.2 (go1.13)
 │    ├── github.com/stretchr/testify@v1.7.0 => [v1.8.2], see <5> for 1 children (branch 'github.com/sirupsen/logrus@v1.9.0' level 2)
 │    ├── golang.org/x/sys@v0.0.0-20210823070655-63515b42dcdf => [v0.8.0]
 │    ├── github.com/davecgh/go-spew@v1.1.0 => [v1.1.1]
 │    ├── github.com/pmezard/go-difflib@v1.0.0
 │    └── gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c => [v3.0.1]
 │         └── 1 more ...
 ├── gocv.io/x/gocv@v0.31.0 (go1.13) => [v0.32.1]
 │    ├── github.com/hybridgroup/mjpeg@v0.0.0-20140228234708-4680f319790e
 │    └── github.com/pascaldekloe/goe@v0.1.0 => [v0.1.1]
 ├── golang.org/x/crypto@v0.1.0 (go1.17) => [v0.9.0]
 │    ├── golang.org/x/net@v0.1.0 (go1.17) => [v0.10.0], see <6> for 3 children (branch 'gobot.io/x/gobot' level 1)
 │    ├── golang.org/x/sys@v0.1.0 (go1.17) => [v0.8.0]
 │    ├── golang.org/x/term@v0.1.0 (go1.17) => [v0.8.0]
 │    └── golang.org/x/text@v0.4.0 (go1.17) => [v0.9.0]
 ├── <6> golang.org/x/net@v0.1.0 (go1.17) => [v0.10.0]
 │    ├── golang.org/x/sys@v0.1.0 (go1.17) => [v0.8.0]
 │    ├── golang.org/x/term@v0.1.0 (go1.17) => [v0.8.0]
 │    └── golang.org/x/text@v0.4.0 (go1.17) => [v0.9.0]
 ├── golang.org/x/sync@v0.1.0 => [v0.2.0]
 ├── golang.org/x/sys@v0.1.0 (go1.17) => [v0.8.0]
 ├── google.golang.org/protobuf@v1.28.1 (go1.11) => [v1.30.0]
 │    ├── github.com/golang/protobuf@v1.5.0 (go1.9) => [v1.5.3], see <7> for 2 children (branch 'gobot.io/x/gobot' level 1)
 │    └── github.com/google/go-cmp@v0.5.5 (go1.8) => [v0.5.9], see <8> for 1 children (branch 'github.com/golang/protobuf@v1.5.0' level 2)
 ├── <3> gopkg.in/yaml.v3@v3.0.1
 │    └── gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405 => [v1.0.0-20201130134442-10cb98267c6c]
 ├── <9> periph.io/x/conn/v3@v3.6.10 (go1.13) => [v3.7.0]
 │    └── github.com/jonboulle/clockwork@v0.2.2 (go1.13) => [v0.4.0]
 ├── periph.io/x/host/v3@v3.7.2 (go1.13) => [v3.8.2]
 │    ├── periph.io/x/conn/v3@v3.6.10 (go1.13) => [v3.7.0], see <9> for 1 children (branch 'gobot.io/x/gobot' level 1)
 │    └── periph.io/x/d2xx@v0.0.4 (go1.13) => [v0.1.0]
 └── tinygo.org/x/bluetooth@v0.6.0 (go1.15)
      ├── github.com/go-ole/go-ole@v1.2.6 (go1.12), see <2> for 1 children (branch 'gobot.io/x/gobot' level 1)
      ├── github.com/godbus/dbus/v5@v5.0.3 => [v5.1.0] (go1.12)
      ├── github.com/muka/go-bluetooth@v0.0.0-20220830075246-0746e3a1ea53 (go1.14) => [v0.0.0-20221213043340-85dc80edc4e1], see <10> for 10 children (branch 'gobot.io/x/gobot' level 1)
      ├── github.com/saltosystems/winrt-go@v0.0.0-20220826130236-ddc8202da421 => [v0.0.0-20230510070731-e096b9afa761]
      │    └── 13 more ...
      ├── github.com/sirupsen/logrus@v1.9.0 (go1.13), see <11> for 3 children (branch 'gobot.io/x/gobot' level 1)
      ├── github.com/tinygo-org/cbgo@v0.0.4 (go1.14), see <12> for 1 children (branch 'gobot.io/x/gobot' level 1)
      ├── golang.org/x/crypto@v0.0.0-20210921155107-089bfa567519 => [v0.9.0]
      │    └── 4 more ...
      ├── golang.org/x/sys@v0.0.0-20220829200755-d48e67d00261 => [v0.8.0]
      ├── tinygo.org/x/drivers@v0.23.0 (go1.15) => [v0.24.0]
      │    └── 6 more ...
      └── tinygo.org/x/tinyterm@v0.1.0 (go1.15) => [v0.2.0]
           └── 4 more ...
```

## Variant: file for dependency graph

Alternatively a file with the content can be provided.

cd into the package root folder (contains go.mod) and run:

`go mod graph > graphfile.txt`

`deptree -graph=graphfile.txt`

## Variant: file for module list

Alternatively a file with the content can be provided

cd into the package root folder (contains go.mod) and run:

`go list -u -m -json all > upgradefile.txt`

`go mod graph | deptree -upgrade=upgradefile.txt`

## Variant: both files used

Relative paths can be used also.

`deptree -graph=data/graphfile.txt -upgrade=data/upgradefile.txt`

## Duplicate occurrence of children

More than one occurrence of an item at different levels of the tree is possible. This occurs very likely for projects
with many dependencies (count of imports in go.mod file). `deptree` handles this in the following way to avoid duplicate
information, waste of processing time and possibly circular dependencies:

* first occurrence of the child will be processed and shown with all sub-children
* already processed children will be filtered by default
* the parameter "-a" will show a link to the first processed item (do not filter the line completely)
* the parameter "-f" forces to process and show each child with all sub-children (stopped at the given depth level)

> The option "-f" can lead to hang due to circular dependencies or great amount of work. If this setting is really
> needed and you note a very long delay, try to set the depth parameter "-d" as minimal as possible.
