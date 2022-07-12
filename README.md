# GoRestApi

This API exposes one endpoint `/noFlyZones`

You can find the api at https://go-faa-restful-api.herokuapp.com/

# /noFlyZones

This endpoint returns the list of No Fly Zones [example](https://go-faa-restful-api.herokuapp.com/noFlyZones)

```json
[
  {
    "spatialReference": {
      "wkid": 102100
    },
    "rings":
    [
      [
        [-9278977.502393615,5196972.662366206],
        [-9278404.224681476,5197240.191965203],
        [-9274505.936238931,5195673.232885358],
        [-9275518.726863708,5190055.1113064],
        [-9278881.956108259,5189061.429938688],
        [-9280869.318843672,5188660.135540191],
        [-9282646.479751302,5192481.986954449],
        [-9278977.502393615,5196972.662366206]
      ]
    ]
  }
]
```
