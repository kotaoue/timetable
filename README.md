# timetable
Create daily pie chart image.

A Go program that creates pie chart images.

## Sample Output

```mermaid
pie
    title Daily Time Allocation
    "Sleep" : 8
    "Work" : 8
    "Leisure" : 4
    "Meals" : 2
    "Exercise" : 1
    "Other" : 1
```

![timetable sample](image.png)

## Requirements

* Go 1.17+
* [Koruri-Regular.ttf](https://koruri.github.io/) placed in the working directory

## Usage

```sh
# Build
go build -o timetable .

# Run with default size (400×400)
./timetable

# Run with custom size
./timetable -w 600 -h 600
```

The output file `image.png` will be created in the current directory.

### Flags

| Flag | Default | Description       |
|------|---------|-------------------|
| `-w` | `400`   | Image width (px)  |
| `-h` | `400`   | Image height (px) |

## Next
Change the implementation plan from self-made to existing.
Next repository is [kotaoue/pieimg](https://github.com/kotaoue/pieimg)

## References
* [Koruri](https://koruri.github.io/)
* [Adobe Color](https://color.adobe.com/ja/create/color-wheel)
* [mermaid](https://mermaid-js.github.io/mermaid/#/)
  * [Theme Configuration](https://mermaid-js.github.io/mermaid/#/theming)