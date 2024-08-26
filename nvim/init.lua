
vim.opt.encoding = "utf-8"
vim.opt.fileformat = "unix"
vim.opt.ruler = true
vim.opt.relativenumber = true
vim.opt.number = true

vim.api.nvim_create_autocmd("FileType", {
  pattern = "makefile",
  command = "setlocal noexpandtab"
})

vim.api.nvim_create_autocmd("FileType", {
  pattern = "md",
  command = "setlocal noexpandtab"
})

vim.opt.foldnestmax = 10
vim.opt.foldenable = false
vim.opt.foldlevel = 1

vim.opt.fillchars:append { vert = "│" }
vim.opt.foldcolumn = "0"
vim.cmd [[
  hi LineNr ctermbg=white ctermfg=black
  hi VertSplit ctermbg=black ctermfg=white
]]

vim.opt.hlsearch = true
vim.opt.ignorecase = true
vim.opt.smartcase = true
vim.opt.showmatch = true
vim.opt.showmode = true
vim.opt.showcmd = true
vim.opt.gdefault = true
vim.opt.cursorline = true
vim.opt.cursorcolumn = true
vim.opt.wrap = true
vim.opt.mouse = "a"
vim.opt.timeoutlen = 1000
vim.opt.ttimeoutlen = 0
vim.opt.modelines = 0
vim.opt.undofile = true
vim.opt.laststatus = 2

-- Remember last edit position
vim.api.nvim_create_autocmd("BufReadPost", {
  pattern = "*",
  callback = function()
    if vim.fn.line("'\"") > 1 and vim.fn.line("'\"") <= vim.fn.line("$") then
      vim.cmd("normal! g'\"")
    end
  end
})

vim.cmd [[
  set nocompatible
  call plug#begin('~/.nvim/plugged')

  Plug 'scrooloose/nerdtree'
  Plug 'Valloric/YouCompleteMe', {'do': './install.py --gocode-completer --tern-completer --rust-completer  --clang-completer'}
  Plug 'rdnetto/YCM-Generator', { 'branch': 'stable'}
  Plug 'mdempsky/gocode', { 'rtp': 'vim', 'do': '~/.vim/plugged/gocode/vim/symlink.sh' }
  Plug 'SirVer/ultisnips'
  Plug 'honza/vim-snippets'
  Plug 'ternjs/tern_for_vim'
  Plug 'pangloss/vim-javascript'
  Plug 'posva/vim-vue'
  Plug 'othree/html5.vim'
  Plug 'rust-lang/rust.vim'
  Plug 'godlygeek/tabular'
  Plug 'plasticboy/vim-markdown'
  Plug 'neomake/neomake'
  Plug 'majutsushi/tagbar'
  Plug 'tomtom/tcomment_vim'
  Plug 'tpope/vim-surround'
  Plug 'tpope/vim-abolish'
  Plug 'mileszs/ack.vim'
  Plug 'airblade/vim-gitgutter'
  Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }
  Plug 'terryma/vim-multiple-cursors'
  Plug 'CodeFalling/fcitx-vim-osx'
  Plug 'bronson/vim-visual-star-search'
  Plug 'raimondi/delimitmate'
  Plug 'ctrlpvim/ctrlp.vim'
  Plug 'ryanoasis/vim-devicons'
  Plug 'editorconfig/editorconfig-vim'
  Plug 'cespare/vim-toml'
  Plug 'buoto/gotests-vim'
  Plug 'chase/vim-ansible-yaml'
  Plug 'altercation/vim-colors-solarized'
  Plug 'sheerun/vim-polyglot'
  Plug 'nvim-treesitter/nvim-treesitter', {'do': ':TSUpdate'}
  Plug 'junegunn/fzf', { 'dir': '~/.fzf', 'do': './install --all' }
  Plug 'nvim-lualine/lualine.nvim'
  Plug 'nvim-tree/nvim-web-devicons'
  Plug 'Exafunction/codeium.vim', { 'branch': 'main' }
  Plug 'rhysd/git-messenger.vim'
  Plug 'tpope/vim-rsi'

  Plug 'ibhagwan/fzf-lua', {'branch': 'main'}
  Plug 'nvim-tree/nvim-web-devicons'

  Plug 'dracula/vim', { 'as': 'dracula' }

  call plug#end()
]]

vim.g.mapleader = " "

-- NERDTree configuration
vim.g.NERDTreeIgnore = {'node_modules$[[dir]]'}
vim.g.NERDTreeShowLineNumbers = 1
vim.api.nvim_create_autocmd("FileType", {
  pattern = "nerdtree",
  command = "setlocal relativenumber"
})

-- YouCompleteMe configuration
-- vim.g.ycm_filetype_blacklist = {}
vim.g.ycm_auto_trigger = 1
vim.g.ycm_server_keep_logfiles = 1
vim.g.ycm_server_log_level = 'debug'
vim.g.ycm_min_num_of_chars_for_completion = 2
vim.g.ycm_key_list_select_completion = {'<c-j>', '<Down>'}
vim.g.ycm_key_list_previous_completion = {'<c-k>', '<Up>'}
vim.g.ycm_key_invoke_completion = '<c-z>'
vim.g.ycm_confirm_extra_conf = 0

-- UltiSnips configuration
vim.g.UltiSnipsExpandTrigger = "<tab>"
vim.g.UltiSnipsJumpForwardTrigger = "<c-l>"
vim.g.UltiSnipsJumpBackwardTrigger = "<c-h>"
vim.g.UltiSnipsEditSplit = "vertical"
vim.g.UltiSnipsSnippetStorageDirectoryForUltiSnipsEdit = '/Users/jialinwu/git/configuration/UltiSnips'
vim.g.UltiSnipsSnippetDirectories = {"UltiSnips", "/Users/jialinwu/git/configuration/UltiSnips"}
vim.g.UltiSnipsSnippetsDir = "/Users/jialinwu/git/configuration/UltiSnips"

-- Auto format for js, html, vue files
vim.api.nvim_create_autocmd("BufWritePre", {
  pattern = {"*.js", "*.html", "*.htm", "*.vue"},
  command = "normal migg=G`izz"
})

-- Set filetype for markdown files
vim.api.nvim_create_autocmd({"BufNewFile", "BufRead"}, {
  pattern = "*.{md,mkd,mkdn,mark*}",
  command = "setlocal filetype=markdown"
})

-- Set filetype for Django files
vim.api.nvim_create_autocmd({"BufNewFile", "BufRead"}, {
  pattern = "*.hj",
  command = "setlocal filetype=htmldjango"
})

-- Neomake configuration
vim.api.nvim_create_autocmd("BufWritePost", {
  pattern = "*",
  command = "Neomake"
})
vim.g.neomake_open_list = 2
vim.g.neomake_go_enabled_makers = { 'go', 'golangci_lint' }
vim.g.neomake_html_enabled_makers = {}
vim.g.neomake_css_enabled_makers = {}

-- Tagbar configuration
vim.g.tagbar_show_linenumbers = 2
vim.g.tagbar_type_markdown = {
  ctagstype = 'markdown',
  kinds = {
    'h:headings',
  },
  sort = 0
}

-- Ack configuration
if vim.fn.executable('rg') == 1 then
  vim.g.ackprg = 'rg --vimgrep'
end

-- GitGutter configuration
vim.g.gitgutter_signs = 1
vim.g.gitgutter_highlight_lines = 0

-- vim-go configuration
vim.g.go_fmt_command = 'goimports'
vim.g.go_def_mode = 'gopls'
vim.g.go_info_mode = 'gopls'

-- Go file type specific mappings
vim.api.nvim_create_autocmd("FileType", {
  pattern = "go",
  callback = function()
    vim.api.nvim_buf_set_keymap(0, 'n', '<Leader>h', '<Plug>(go-def-split)', {noremap = true, silent = true})
    vim.api.nvim_buf_set_keymap(0, 'n', '<Leader>v', '<Plug>(go-def-vertical)', {noremap = true, silent = true})
    vim.api.nvim_buf_set_keymap(0, 'n', '<Leader>t', '<Plug>(go-def-tab)', {noremap = true, silent = true})
  end
})

-- Highlight settings
vim.cmd [[
  highlight Pmenu ctermfg=254 ctermbg=33
  highlight PmenuSel ctermfg=235 ctermbg=39 cterm=bold
]]

-- YAML file settings
vim.api.nvim_create_autocmd({"BufNewFile", "BufReadPost"}, {
  pattern = "*.{yaml,yml}",
  command = "set filetype=yaml foldmethod=indent"
})
vim.api.nvim_create_autocmd("FileType", {
  pattern = "yaml",
  command = "setlocal ts=2 sts=2 sw=2 expandtab"
})

-- Color scheme settings
if vim.fn.has('termguicolors') == 1 then
  vim.opt.termguicolors = true
end
vim.opt.background = "light"
vim.cmd [[colorscheme solarized8]]
-- vim.opt.background = "dark"
-- vim.cmd [[colorscheme dracula]]


-- Custom variables
vim.g.vimrc_author = 'Jialin Wu'
vim.g.vimrc_email = 'win27v@gmail.com'
vim.g.vimrc_homepage = 'https://jialinwu.com'

-- FZF configuration
vim.opt.rtp:append("/Users/jialinwu/.fzf")

-- Codeium configuration
vim.g.codeium_disable_bindings = 1
vim.keymap.set('i', '<C-g>', function () return vim.fn['codeium#Accept']() end, { expr = true, silent = true })
vim.keymap.set('i', '<c-;>', function() return vim.fn['codeium#CycleCompletions'](1) end, { expr = true, silent = true })
vim.keymap.set('i', '<c-,>', function() return vim.fn['codeium#CycleCompletions'](-1) end, { expr = true, silent = true })
vim.keymap.set('i', '<c-x>', function() return vim.fn['codeium#Clear']() end, { expr = true, silent = true })

-- Custom keymappings
vim.api.nvim_set_keymap('n', '<leader>r', ':CtrlPMRU<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>f', ':FZF<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>nt', ':NERDTreeToggle<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>e', ':NERDTreeToggle<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>q', ':q!<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>1q', ':qa!<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>x', ':x<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>a', 'ggvGV<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>qa', ':qa!<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>b', ':NERDTreeFromBookmark ', {noremap = true})

vim.api.nvim_set_keymap('n', '<leader>d', '"+d', {noremap = true})
vim.api.nvim_set_keymap('n', '<leader>y', '"+y', {noremap = true})
vim.api.nvim_set_keymap('n', '<leader>1y', '"1y', {noremap = true})
vim.api.nvim_set_keymap('n', '<leader>2y', '"2y', {noremap = true})

vim.api.nvim_set_keymap('n', '<leader>p', '"+p', {noremap = true})
vim.api.nvim_set_keymap('n', '<leader>P', '"+P', {noremap = true})
vim.api.nvim_set_keymap('n', '<leader>1p', '"1p', {noremap = true})
vim.api.nvim_set_keymap('n', '<leader>1P', '"1P', {noremap = true})
vim.api.nvim_set_keymap('n', '<leader>2p', '"2p', {noremap = true})
vim.api.nvim_set_keymap('n', '<leader>2P', '"2P', {noremap = true})

vim.api.nvim_set_keymap('n', '<leader>af', ':echo expand("%:p")<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>k', ':Ack ', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>tn', ':tabnew<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>tc', ':tabclose<cr>', {noremap = true, silent = true})

vim.api.nvim_set_keymap('t', '<C-\\>', '<C-\\><C-n>', {noremap = true})
vim.api.nvim_set_keymap('n', '<C-h>', '<C-w>h', {noremap = true})
vim.api.nvim_set_keymap('n', '<C-j>', '<C-w>j', {noremap = true})
vim.api.nvim_set_keymap('n', '<C-k>', '<C-w>k', {noremap = true})
vim.api.nvim_set_keymap('n', '<C-l>', '<C-w>l', {noremap = true})

vim.api.nvim_set_keymap('n', ';', ':', {noremap = true})
vim.api.nvim_set_keymap('n', 'w=', ':resize +20<CR>', {noremap = true})
vim.api.nvim_set_keymap('n', 'w-', ':resize -20<CR>', {noremap = true})
vim.api.nvim_set_keymap('n', 'w,', ':vertical resize -20<CR>', {noremap = true})
vim.api.nvim_set_keymap('n', 'w.', ':vertical resize +20<CR>', {noremap = true})

-- For C++/C
vim.api.nvim_set_keymap('n', '<leader>hc', ':e %:p:s,.h$,.X123X,:s,.cpp$,.h,:s,.X123X$,.cpp,<cr>', {noremap = true, silent = true})

-- Edit tmp file
vim.api.nvim_set_keymap('n', '<leader>tmp', ':e ~/tmpfile<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>booknote', ':e /Users/Jialin/Dropbox/vimNote/book.md<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>jsnote', ':e /Users/Jialin/Dropbox/vimNote/jsnote.md<cr>', {noremap = true, silent = true})

-- Highlight and search
vim.api.nvim_set_keymap('v', '//', 'y/<C-R>"<CR>', {noremap = true})

-- Make
vim.api.nvim_set_keymap('n', '<leader>zm', ':make <cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>zl', ':lopen<cr>', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '<leader>m', ':TagbarToggle<cr>', {noremap = true, silent = true})

-- Tab to %
vim.api.nvim_set_keymap('n', '<tab>', '%', {noremap = true})
vim.api.nvim_set_keymap('v', '<tab>', '%', {noremap = true})

-- Keep highlighting while doing indent
vim.api.nvim_set_keymap('v', '<', '<gv', {noremap = true})
vim.api.nvim_set_keymap('v', '>', '>gv', {noremap = true})

-- Select text from current cursor to the end of line
vim.api.nvim_set_keymap('n', 'Y', 'v$', {noremap = true})

-- Delete text from current cursor to the end or beginning of line in insert mode
vim.api.nvim_set_keymap('!', '<M-l>', '<ESC>v$c', {noremap = true})
vim.api.nvim_set_keymap('!', '<M-h>', '<ESC>v^c', {noremap = true})

-- Keep search pattern at the center of the screen
vim.api.nvim_set_keymap('n', 'n', 'nzz', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', 'N', 'Nzz', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '*', '*zz', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', '#', '#zz', {noremap = true, silent = true})
vim.api.nvim_set_keymap('n', 'g*', 'g*zz', {noremap = true, silent = true})

-- Better command-line editing
vim.api.nvim_set_keymap('c', '<C-a>', '<Home>', {noremap = true})
vim.api.nvim_set_keymap('c', '<C-e>', '<End>', {noremap = true})

-- Replace : with ; in normal mode
vim.api.nvim_set_keymap('n', ';', ':', {noremap = true})
vim.api.nvim_set_keymap('n', "'", ';', {noremap = true})
vim.api.nvim_set_keymap('n', '"', ',', {noremap = true})

-- Split window
vim.api.nvim_set_keymap('n', '<leader>2w', '<C-w>v<C-w>l', {noremap = true})

-- Visual for HTML tag
vim.api.nvim_set_keymap('n', '<leader>zv', 'Vat', {noremap = true})

-- Use alt-w to save
vim.api.nvim_set_keymap('n', '<M-w>', ':w<cr>', {noremap = true})
vim.api.nvim_set_keymap('i', '<M-w>', '<ESC>:w<cr>', {noremap = true})

vim.api.nvim_set_keymap('i', 'lk', '<ESC>', {noremap = true})

-- Open dir of nerd-tree on current file
vim.api.nvim_set_keymap('n', '<Leader>cd', ':NERDTree %:p:h<CR>', {noremap = true})

-- Save file editing
vim.api.nvim_set_keymap('n', '<leader>x', ':w<cr>', {noremap = true, silent = true})

-- Fast editing of editor configuration
vim.api.nvim_set_keymap('n', '<leader>ee', ':e ~/.config/nvim/init.lua<cr>', {noremap = true, silent = true})

-- Auto-reload configuration on save
vim.cmd [[
  augroup config_reload
    autocmd!
    autocmd BufWritePost init.lua source ~/.config/nvim/init.lua
  augroup END
]]

-- Set highlight for Directory
vim.cmd [[hi Directory ctermfg=Blue]]

-- Set Python and Ctags paths
vim.g.python3_host_prog = '/opt/homebrew/bin/python3'
vim.g.tagbar_ctags_bin = '/opt/homebrew/bin/ctags'

-- Additional settings from the original vimrc
vim.g.go_fmt_command = 'goimports'
vim.g.go_def_mode = 'gopls'
vim.g.go_info_mode = 'gopls'

-- -- Set up nvim-treesitter
-- require'nvim-treesitter.configs'.setup {
--   ensure_installed = "all",
--   highlight = {
--     enable = true,
--   },
-- }

-- Set up statusline
vim.opt.statusline = vim.opt.statusline + [[{…}%3{codeium#GetStatusString()}]]

-- Set up custom functions (you may need to adjust these for Lua)
_G.LightLineFilename = function()
  return vim.fn.expand('%')
end

-- Sync clipboard between OS and Neovim.
--  Schedule the setting after `UiEnter` because it can increase startup-time.
--  Remove this option if you want your OS clipboard to remain independent.
--  See `:help 'clipboard'`
vim.schedule(function()
  vim.opt.clipboard = 'unnamedplus'
end)

require("nvim-treesitter.install").prefer_git = true
require('nvim-treesitter.configs').setup({
    ensure_installed = {
				-- 'rust',
				'javascript',
				'html',
				'go',
				'vue',
				'lua',
				'css',
				'cpp',
				'c',
				'json',
				'query', -- for playground
			}
    })

local fzf = require('fzf-lua')

-- Define your workspaces
local workspaces = {
  '~/go/src/SmartOutCall',
  '~/go/src/OutCallBinLog',
  '~/go/src/MessageProxy',
  '~/go/src/CowDialog',
  '~/go/src/AudioProxy',
  '~/go/src/helper',
  '~/git/configuration/note'
}

local base_dir = '~/go/src/1'  -- Directory to include its first-level directories

-- Add first-level directories from base_dir to the workspaces list
local dirs = vim.fn.globpath(base_dir, '*', 0, 1)
for _, dir in ipairs(dirs) do
  table.insert(workspaces, dir)
end

-- Function to switch workspace
function _G.switch_workspace()
  fzf.fzf_exec(workspaces, {
    actions = {
      ['default'] = function(selected)
        -- Expand the path (to handle '~')
        local expanded_path = vim.fn.expand(selected[1])
        -- Change the current working directory
        vim.cmd('cd ' .. expanded_path)
        print('Switched to workspace: ' .. expanded_path)
        -- Show files in the selected workspace using fzf-lua
        fzf.files({ cwd = expanded_path })
      end
    }
  })
end

-- status line
-- solarized_jialinwu config for lualine
-- Author: JialinWu
-- MIT license, see LICENSE for more details.

-- stylua: ignore
local colors = {
  blue = "#5a80f3",
  black  = '#080808',
  white  = '#fdf6e3',
  red    = '#cb4b16',
  green='#489926',
  orange = '#fe8018',
  grey='#839496'

}


local solarized_light_jialinwu_theme = {
  normal = {
    a = { fg = colors.white, bg = colors.green },
    b = { fg = colors.white, bg = colors.blue },
    c = { fg = colors.white },
  },

  insert = { a = { fg = colors.white, bg = colors.red } },
  visual = { a = { fg = colors.black, bg = colors.green } },
  replace = { a = { fg = colors.black, bg = colors.red } },

  inactive = {
    a = { fg = colors.white, bg = colors.green },
    b = { fg = colors.white, bg = colors.grey },
    c = { fg = colors.white },
  },
}

require('lualine').setup {
  options = {
    theme = solarized_light_jialinwu_theme,
    -- theme = 'solarized_light',
    component_separators = '',
    section_separators = { left = '', right = '' },
  },
  sections = {
    lualine_a = { { 'mode', separator = { left = '' }, right_padding = 2 } },
    -- lualine_b = {'branch' ,'filename'},
    lualine_b = {
      'branch',
      'diff',
      {
        'diagnostics',
        source = { 'nvim' },
        sections = { 'error' },
        diagnostics_color = { error = { bg = colors.red, fg = colors.white } },
      },
      {
        'diagnostics',
        source = { 'nvim' },
        sections = { 'warn' },
        diagnostics_color = { warn = { bg = colors.orange, fg = colors.white } },
      },
      { 'filename', file_status = false, path = 1 },
      { modified, color = { bg = colors.red } },
    },
    lualine_c = {
      '%=', --[[ add your center compoentnts here in place of this comment ]]
    },
    lualine_x = {},
    lualine_y = { 'searchcount', 'progress','encoding' },
    lualine_z = {
      { 'location', separator = { right = '' }, left_padding = 2 },
    },
  },
  inactive_sections = {
    lualine_a = {},
    lualine_b = {'branch' ,'filename'},
    lualine_c = {},
    lualine_x = {},
    lualine_y = {},
    lualine_z = { 'location' },
  },
  tabline = {},
  extensions = {},
}

-- Set up a keybinding to trigger the workspace switcher
vim.api.nvim_set_keymap('n', '<leader>w', [[<cmd>lua switch_workspace()<CR>]], { noremap = true, silent = true })

-- End of init.lua
