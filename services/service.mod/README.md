// mod_service.ModService.GetModByID
```sh
grpcurl -plaintext -d '{"ID":"c42f6b43-628a-4aa8-8a2c-3ddbfe7f9f88"}' localhost:4089 mod_service.ModService.GetModByID
```

// mod_service.ModService.SearchMod
```sh
grpcurl -plaintext localhost:4089 mod_service.ModService.SearchMod
```

// mod_service.ModService.SearchMod
```sh
grpcurl -plaintext -d '{"SearchText":"Ultra Violet"}' localhost:4089 mod_service.ModService.SearchMod
```

// mod_service.ModService.SearchMod
```sh
grpcurl -plaintext -d '{"ModTypeCategoryID":"63964035-51ef-40a1-87e3-fa533bbbebc4", "Page":1, "Size":2}' localhost:4089 mod_service.ModService.SearchMod
```

grpcurl -plaintext -d '{"ModTypeCategoryIDs": ["f0cd00f6-d0cc-4247-a0a9-2fcce2bb87bf"]}' localhost:4089 mod_service.ModService.SearchMod

