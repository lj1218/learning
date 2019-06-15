package main

import "sync"

type Image struct {
	name string
}

var icons map[string]Image

func loadIcon(name string) Image {
	return Image{name: name}
}

func loadIcons() {
	icons = map[string]Image{
		"1.png": loadIcon("1.png"),
		"2.png": loadIcon("2.png"),
		"3.png": loadIcon("3.png"),
	}
}

// NOTE: not concurrency-safe!
func IconV1(name string) Image {
	if icons == nil {
		loadIcons() // one-time initialization
	}
	return icons[name]
}

var mu sync.Mutex // guards icons

// Concurrency-safe.
func IconV2(name string) Image {
	mu.Lock()
	defer mu.Unlock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

var mu2 sync.RWMutex // guards icons

// Concurrency-safe.
func IconV3(name string) Image {
	mu2.RLock()
	if icons != nil {
		icon := icons[name]
		mu2.RUnlock()
		return icon
	}
	mu2.RUnlock()

	mu2.Lock()
	if icons == nil { // NOTE: must recheck for nil
		loadIcons()
	}
	icon := icons[name]
	mu2.Unlock()
	return icon
}

var loadIconsOnce sync.Once

// Concurrency-safe.
func IconV4(name string) Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

func main() {}
