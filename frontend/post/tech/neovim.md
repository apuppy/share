# neovim
- [neovim](#neovim)
  - [install](#install)
  - [配置](#配置)
    - [配置文件](#配置文件)
    - [常用快捷键](#常用快捷键)
    - [查看keymap](#查看keymap)
    - [Mason](#mason)
    - [self-maintained kickstart.nvim](#self-maintained-kickstartnvim)

## install
```bash
sudo apt-get install neovim
```

## 配置

### 配置文件
```bash
mkdir -p ~/.config/nvim

# https://github.com/nvim-lua/kickstart.nvim/blob/master/init.lua
edit ~/.config/nvim/init.lua

# should clone whole git repo to ~/.config/nvim
plugin help renamed to vimdoc
```

### 常用快捷键
```bash
- space sf # search file
- gd # go to defination
```

### 查看keymap
```bash
:Telescope keymaps
```

### Mason
```bash
:Mason
```

### self-maintained kickstart.nvim
```bash
git clone https://github.com/apuppy/kickstart.nvim.git ~/.config/nvim
```