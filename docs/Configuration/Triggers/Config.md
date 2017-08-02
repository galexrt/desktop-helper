# Config

## IPAddress
### Options
| Name | Type | Default | Description |
| ---- | ---- | ------- | ----------- |
| Interfaces | `[]string` | "" | List of interfaces to get data from. |

### Examples
```yaml
ipaddress:
  interfaces:
  - eth0
  - wlan0
```

## Xrandr
### Options
| Name | Type | Default | Description |
| ---- | ---- | ------- | ----------- |
| XrandrBinary | `string` | `` | Path to the Xrandr binary. |
| Xauthoritiy | `[]string` | `` | List of interfaces to get data from. |
| IgnoreSegFault | `bool` | `true` | If Segfaults from Xrandr should be ignored. |
| IgnoreErrors | `bool` | `` | Ignore all errors from Xrandr. |
| Display | `string` | `` | Which X `DISPLAY` to use. |
| ScreensIgnore | `[]string` | `` Which screens should be ignored and not listed. |

### Examples
```yaml
xrandr:
  xauthoritiy: /home/galexrt/.Xauthoritiy
  xrandrbinary: /usr/bin/xrandr
  ignoresegfault: true
  ignoreerrors: true
  display: ":0"
```
