# deptree

show golang dependence like tree

## How to get it

```shell
➜ go install github.com/vc60er/deptree@latest
➜ export PATH=$PATH:$GOPATH/bin
```

## Usage

### params

```text
Usage of ./output/deptree:
  -d int
      max depth of dependencies (default 3)
  -graph string
      path to file created e.g. by 'go mod graph > grapphfile.txt'
  -t  visualize trimmed tree by '└─...'
```

### example

```shell
➜  redis git:(master) go mod graph | deptree -d 2 -t
dependency tree with depth 2 for package: gobot.io/x/gobot, least 622 trimmed item(s)

gobot.io/x/gobot
 ├── github.com/bmizerany/pat@v0.0.0-20210406213842-e4b6760bdd6f
 ├── github.com/cpuguy83/go-md2man/v2@v2.0.2
 │    └── github.com/russross/blackfriday/v2@v2.1.0
 ├── github.com/creack/goselect@v0.1.2
 ├── github.com/davecgh/go-spew@v1.1.1
 ├── github.com/donovanhide/eventsource@v0.0.0-20210830082556-c59027999da0
 ├── github.com/eclipse/paho.mqtt.golang@v1.4.1
 │    ├── github.com/gorilla/websocket@v1.4.2
 │    ├── golang.org/x/net@v0.0.0-20200425230154-ff2c4b7c35a0
 │    │    └── ...
 │    └── golang.org/x/sync@v0.0.0-20210220032951-036812b2e83c
 ├── github.com/fatih/structs@v1.1.0
 ├── github.com/go-ole/go-ole@v1.2.6
 │    └── golang.org/x/sys@v0.0.0-20190916202348-b4ddaad3f8a3
 ├── github.com/godbus/dbus/v5@v5.1.0
 ├── github.com/gofrs/uuid@v4.3.0+incompatible
 ├── github.com/golang/protobuf@v1.5.0
 │    ├── github.com/google/go-cmp@v0.5.5
 │    │    └── ...
 │    └── google.golang.org/protobuf@v1.26.0-rc.1
 │         └── ...
 ├── github.com/gorilla/websocket@v1.5.0
 ├── github.com/hashicorp/errwrap@v1.1.0
 ├── github.com/hashicorp/go-multierror@v1.1.1
 │    └── github.com/hashicorp/errwrap@v1.0.0
 ├── github.com/hybridgroup/go-ardrone@v0.0.0-20140402002621-b9750d8d7b78
 ├── github.com/hybridgroup/mjpeg@v0.0.0-20140228234708-4680f319790e
 ├── github.com/muka/go-bluetooth@v0.0.0-20220830075246-0746e3a1ea53
 │    ├── github.com/fatih/structs@v1.1.0
 │    ├── github.com/godbus/dbus/v5@v5.0.3
 │    ├── github.com/niemeyer/pretty@v0.0.0-20200227124842-a10e7caefd8e
 │    │    └── ...
 │    ├── github.com/pkg/errors@v0.9.1
 │    ├── github.com/sirupsen/logrus@v1.6.0
 │    │    └── ...
 │    ├── github.com/stretchr/testify@v1.6.1
 │    │    └── ...
 │    ├── github.com/suapapa/go_eddystone@v1.3.1
 │    │    └── ...
 │    ├── golang.org/x/sys@v0.0.0-20200728102440-3e129f6d46b1
 │    ├── golang.org/x/tools@v0.0.0-20200925191224-5d1fdd8fa346
 │    │    └── ...
 │    └── gopkg.in/check.v1@v1.0.0-20200227125254-8fa46927fb4f
 ├── github.com/nats-io/nats-server/v2@v2.1.0
 │    ├── github.com/golang/protobuf@v1.3.2
 │    ├── github.com/nats-io/jwt@v0.3.0
 │    │    └── ...
 │    ├── github.com/nats-io/nats.go@v1.8.1
 │    │    └── ...
 │    ├── github.com/nats-io/nkeys@v0.1.0
 │    │    └── ...
 │    ├── github.com/nats-io/nuid@v1.0.1
 │    ├── golang.org/x/crypto@v0.0.0-20190701094942-4def268fd1a4
 │    │    └── ...
 │    └── golang.org/x/sys@v0.0.0-20190726091711-fc99dfbffb4e
 ├── github.com/nats-io/nats.go@v1.18.0
 │    ├── github.com/nats-io/nkeys@v0.3.0
 │    │    └── ...
 │    └── github.com/nats-io/nuid@v1.0.1
 ├── github.com/nats-io/nkeys@v0.3.0
 │    └── golang.org/x/crypto@v0.0.0-20210314154223-e6e6c4f2bb5b
 │         └── ...
 ├── github.com/nats-io/nuid@v1.0.1
 ├── github.com/pkg/errors@v0.9.1
 ├── github.com/pmezard/go-difflib@v1.0.0
 ├── github.com/russross/blackfriday/v2@v2.1.0
 ├── github.com/saltosystems/winrt-go@v0.0.0-20220913104103-712830fcd2ad
 │    ├── github.com/glerchundi/subcommands@v0.0.0-20181212083838-923a6ccb11f8
 │    ├── github.com/go-kit/log@v0.2.1
 │    │    └── ...
 │    ├── github.com/go-ole/go-ole@v1.2.6
 │    │    └── ...
 │    ├── github.com/peterbourgon/ff/v3@v3.1.2
 │    │    └── ...
 │    ├── github.com/stretchr/testify@v1.7.5
 │    │    └── ...
 │    ├── github.com/tdakkota/win32metadata@v0.1.0
 │    │    └── ...
 │    ├── golang.org/x/tools@v0.1.11
 │    │    └── ...
 │    ├── github.com/davecgh/go-spew@v1.1.1
 │    ├── github.com/go-logfmt/logfmt@v0.5.1
 │    ├── github.com/pmezard/go-difflib@v1.0.0
 │    ├── golang.org/x/mod@v0.6.0-dev.0.20220419223038-86c51ed26bb4
 │    │    └── ...
 │    ├── golang.org/x/sys@v0.0.0-20220624220833-87e55d714810
 │    └── gopkg.in/yaml.v3@v3.0.1
 │         └── ...
 ├── github.com/sigurn/crc8@v0.0.0-20220107193325-2243fe600f9f
 ├── github.com/sirupsen/logrus@v1.9.0
 │    ├── github.com/davecgh/go-spew@v1.1.1
 │    ├── github.com/stretchr/testify@v1.7.0
 │    │    └── ...
 │    └── golang.org/x/sys@v0.0.0-20220715151400-c0bba94af5f8
 ├── github.com/stretchr/testify@v1.8.0
 │    ├── github.com/davecgh/go-spew@v1.1.1
 │    ├── github.com/pmezard/go-difflib@v1.0.0
 │    ├── github.com/stretchr/objx@v0.4.0
 │    │    └── ...
 │    └── gopkg.in/yaml.v3@v3.0.1
 │         └── ...
 ├── github.com/tinygo-org/cbgo@v0.0.4
 │    └── github.com/sirupsen/logrus@v1.5.0
 │         └── ...
 ├── github.com/urfave/cli@v1.22.10
 │    ├── github.com/BurntSushi/toml@v0.3.1
 │    ├── github.com/cpuguy83/go-md2man/v2@v2.0.0-20190314233015-f79a8a8ca69d
 │    │    └── ...
 │    └── gopkg.in/yaml.v2@v2.2.2
 │         └── ...
 ├── github.com/veandco/go-sdl2@v0.4.25
 ├── github.com/warthog618/gpiod@v0.8.0
 │    ├── github.com/pilebones/go-udev@v0.0.0-20180820235104-043677e09b13
 │    ├── github.com/spf13/cobra@v0.0.5
 │    ├── github.com/spf13/pflag@v1.0.5
 │    ├── github.com/stretchr/testify@v1.4.0
 │    ├── github.com/warthog618/config@v0.4.1
 │    ├── golang.org/x/sys@v0.0.0-20200223170610-d5e6a3e2c0ae
 │    ├── gopkg.in/check.v1@v1.0.0-20190902080502-41f04d3bba15
 │    ├── github.com/davecgh/go-spew@v1.1.1
 │    ├── github.com/fsnotify/fsnotify@v1.4.7
 │    ├── github.com/inconshreveable/mousetrap@v1.0.0
 │    ├── github.com/pkg/errors@v0.8.1
 │    ├── github.com/pmezard/go-difflib@v1.0.0
 │    └── gopkg.in/yaml.v2@v2.2.2
 │         └── ...
 ├── go.bug.st/serial@v1.4.0
 │    ├── github.com/creack/goselect@v0.1.2
 │    ├── github.com/stretchr/testify@v1.7.0
 │    │    └── ...
 │    ├── golang.org/x/sys@v0.0.0-20210823070655-63515b42dcdf
 │    ├── github.com/davecgh/go-spew@v1.1.0
 │    ├── github.com/pmezard/go-difflib@v1.0.0
 │    └── gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c
 │         └── ...
 ├── gocv.io/x/gocv@v0.31.0
 │    ├── github.com/hybridgroup/mjpeg@v0.0.0-20140228234708-4680f319790e
 │    └── github.com/pascaldekloe/goe@v0.1.0
 ├── golang.org/x/crypto@v0.1.0
 │    ├── golang.org/x/net@v0.1.0
 │    │    └── ...
 │    ├── golang.org/x/sys@v0.1.0
 │    ├── golang.org/x/term@v0.1.0
 │    └── golang.org/x/text@v0.4.0
 ├── golang.org/x/net@v0.1.0
 │    ├── golang.org/x/sys@v0.1.0
 │    ├── golang.org/x/term@v0.1.0
 │    └── golang.org/x/text@v0.4.0
 ├── golang.org/x/sync@v0.1.0
 ├── golang.org/x/sys@v0.1.0
 ├── google.golang.org/protobuf@v1.28.1
 │    ├── github.com/golang/protobuf@v1.5.0
 │    │    └── ...
 │    └── github.com/google/go-cmp@v0.5.5
 │         └── ...
 ├── gopkg.in/yaml.v3@v3.0.1
 │    └── gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
 ├── periph.io/x/conn/v3@v3.6.10
 │    └── github.com/jonboulle/clockwork@v0.2.2
 ├── periph.io/x/host/v3@v3.7.2
 │    ├── periph.io/x/conn/v3@v3.6.10
 │    │    └── ...
 │    └── periph.io/x/d2xx@v0.0.4
 └── tinygo.org/x/bluetooth@v0.6.0
      ├── github.com/go-ole/go-ole@v1.2.6
      │    └── ...
      ├── github.com/godbus/dbus/v5@v5.0.3
      ├── github.com/muka/go-bluetooth@v0.0.0-20220830075246-0746e3a1ea53
      │    └── ...
      ├── github.com/saltosystems/winrt-go@v0.0.0-20220826130236-ddc8202da421
      │    └── ...
      ├── github.com/sirupsen/logrus@v1.9.0
      │    └── ...
      ├── github.com/tinygo-org/cbgo@v0.0.4
      │    └── ...
      ├── golang.org/x/crypto@v0.0.0-20210921155107-089bfa567519
      │    └── ...
      ├── golang.org/x/sys@v0.0.0-20220829200755-d48e67d00261
      ├── tinygo.org/x/drivers@v0.23.0
      │    └── ...
      └── tinygo.org/x/tinyterm@v0.1.0
           └── ...

```
