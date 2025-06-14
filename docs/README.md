An online shop application, mono-repo microservices demo for kratos.

[Design](design.md)

## Kratos Mono-Repo structure
```
.
├── api  // API&Error Proto files & Generated codes
│   ├── foo
│   │   ├── job
│   │   └── service
│   └── bar
│       └── interface
├── app  // kratos microservices projects
│   ├── foo
│   │   ├── job
│   │   └── service
│   └── bar
│       └── interface
├── pkg  // common used packages
└── docs

```
