# Mapture
Mapture is a CLI + TUI tool that parses your application config files (like Aerospace, Neovim, and more) and lets you list and fuzzy‑search your keymaps across tools — with pretty output and an interactive table view.

Because life is too short to grep through dotfiles.


---
## 🚀  Features

- Knows how to parse popular config files (Aerospace, Neovim, … more coming like Tmux, i3).
- Unified Keymap model (Source, Mode, Shortcut, Actions).
- Fuzzy search across all your keymaps (via fuzzysearch).
- TUI powered by Bubble Tea + Bubbles/Table:
    - Live Filter as you type.

---
### 💡  Usage Examples

#### 📃  List keymaps

```bash
# List everything across all known config sources
mapture list

# Limit to Aerospace only
mapture list --source aerospace

# Limit to Nvim only
mapture list --source nvim --file ~/.config/nvim/init.lua

```
##### Sample Output
```
alt-f        → fullscreen          [aerospace] (main)
<leader>ff   → Telescope find_files [nvim]     (normal)
```

#### 🔎  Search keymaps
```bash 
# Search across all sources (default)
mapture search fullscreen

# Search only in Nvim
mapture search telescope --source nvim

```

##### Sample Output 

```
<leader>ff   → Telescope find_files [nvim]     (normal)
```

🎨 ## Interactive TUI Search
```bash
mapture search -tui 
```
#### Features

- Type to filter results live.
- Navigate with ↑/↓ or j/k.
- Press Enter to copy the selected mapping to your clipboard.
- Exit with Esc or Ctrl+C.

---

## 🛣️  Roadmap
 - [ ] Add Tmux parser (default ~/.tmux.conf).
 - [ ] Add i3/Sway parser (~/.config/i3/config).
 - [ ]  Add mapture export --format markdown to dump cheatsheets.
 - [ ] Add mapture index to auto‑scan ~/.config/ and build a global keymap index.
 - [ ] Add column sorting in TUI (press S / K etc.)

## 🔧 Installation

### Build from source
```bash 
git clone https://github.com/thomasmclean993/mapture
cd mapture
go build -o mapture 
```

### Run Locally 
```bash 
./mapture list
./mapture search fullscreen
./mapture search --tui 
```
