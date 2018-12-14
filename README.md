

## Firehose Chart

Given an array of metrics this application will spawn a terminal application using [termui](https://github.com/gizak/termui.git) that creates a line chart for the given metrifcs


Example below will graph the user cpu usage for a single doppler instance 

```
firehose-chart -m '[{ "origin": "bosh-system-metrics-forwarder", "type": "ValueMetric", "job": "doppler", "index": "fa92e017-b23f-4fdc-8630-6d6080857d83", "metric": "system.cpu.user"} ]'
```

Example interface:

```
┌Monitoring Firehose───────────────────────────────────────────────────────────┐
│     ┊                                                                        │
│40.84┊                                                                        │
│     ┊                             ⠂ ⠂           ⠐ ⠐              ⠄           │
│31.29┊                                                          ⠁             │
│     ┊                                                                        │
│21.75┊                                                              ⠄         │
│     ┊                           ⠄                                            │
│12.21┊                                         ⢀     ⢀                        │
│     ┊                                                                        │
│2.66 ┊                      ⠒⠒⠒⠒⠒⠐⠒⠐⠒⠐⠒⠒⠒⠒⠒⠒⠒⠒⠒⠂⠒⠂⠒⠂⠒⠂⠒⠒⠒⠒⠒⠒⠒⠒⠒⠒⠐⠒⠐⠒⠐⠒⠒⠒⠒⠒⠒⠒⠒⠒│
│     ┊                                                                        │
│-6.88└┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈┈ │
│     0  6  12  20  28  36  44  52  60  68  76  84  92                         │
└──────────────────────────────────────────────────────────────────────────────┘

┌legend────────────────────────────────────────────────────────────────────────┐
│doppler/fa92e017-b23f-4fdc-8630-6d6080857d83 = yellow                         │
│                                                                              │
│                                                                              │
└──────────────────────────────────────────────────────────────────────────────┘
```















