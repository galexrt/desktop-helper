# Profile Triggers
Triggers that cause a profile to be matched, and enabled and disabled.

## IPAddress
### Options
Name | Type | Default | Description
-----------------------------------
Addresses | `map[string]IPAddress` | "" | List.

#### IPAddress
Name | Type | Default | Description
-----------------------------------
Address | `string` | "" | The actual IP address to be equal with the subnet cidr (`192.168.1.10/24`).
Key | `string` | "" | The "key" of the ip in order shown by `ip addr show`.

### Examples
```yaml
ipaddress:
  addresses:
    eth0:
      address: 192.168.1.10/24
      key: 0
```

## Xrandr
### Options
Name | Type | Default | Description
-----------------------------------
ConnectedCount | `int` | "" | The count of screens connected.
Screens | `[]string` | "" | List of screens that should be connected.

### Examples
```yaml
xrandr:
  connectedcount: 2
  screens:
    - eDP1
    - DP1-1
```
