" cat > ~/.config/nvim/init.vim
set encoding=utf-8
set ff=unix
set ruler "display the current row and column on the right_bottom corner 
set relativenumber
set nu

" 空格代替Tab 
" set expandtab
" 缩进宽度
" set tabstop=4 
" set shiftwidth=4
" 禁在Makefile 中将Tab转换成空格
autocmd FileType makefile set noexpandtab
autocmd FileType md set noexpandtab

" set foldmethod=syntax "fold based on indent
set foldnestmax=10      "deepest fold is 10 levels
set nofoldenable        "dont fold by default
set foldlevel=1         "this is just what i use

set fillchars+=vert:│
" set fillchars+=vert:\ 
set foldcolumn=0
hi LineNr ctermbg=white ctermfg=black
" hi LineNr ctermbg=white ctermfg=red
hi VertSplit ctermbg=black ctermfg=white

set hlsearch
" set backspace=indent,eol,start

" ignorecase and smartcase together make Vim deal with case-sensitive search intelligently. If you search for an all-lowercase string your search will be case-insensitive, but if one or more characters is uppercase the search will be case-sensitive. Most of the time this does what you want.
set ignorecase
set smartcase
set showmatch
set showmode
set showcmd
set gdefault

set cursorline
set cursorcolumn

" set colorcolumn=85
set wrap
set mouse=a

set timeoutlen=1000 ttimeoutlen=0

set modelines=0   "don't execute command that comment in a file"
set undofile  "this allow you to undo change even when you reopen a file
set laststatus=2   " Always show the statusline

"remember last edicting position
au BufReadPost * if line("'\"") > 0|if line("'\"") <= line("$")|exe("norm '\"")|else|exe "norm $"|endif|endif

" https://github.com/sainnhe/edge?tab=readme-ov-file
" ====================plugin manager=================
set nocompatible
call plug#begin('~/.nvim/plugged')
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'scrooloose/nerdtree'
let NERDTreeIgnore=['node_modules$[[dir]]']
" enable line numbers
let NERDTreeShowLineNumbers=1
" make sure relative line numbers are used
autocmd FileType nerdtree setlocal relativenumber
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""<code completion go-vim> start""""""""""""""""""""""""""""""""""
"conflict with ycm 2019-12-16 11:00:44
" Plug 'govim/govim'
"""""""""""""""""""""""""""""<code completion go-vim> end""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""<code completion> start""""""""""""""""""""""""""""""""""
" Plug 'Valloric/YouCompleteMe',{'do': './install.py --clang-completer --tern-completer'}
"  should use g:python3_host_prog e.g. /opt/homebrew/bin/python3 ./install.py --gocode-completer --tern-completer --rust-completer  --clang-completer
Plug 'Valloric/YouCompleteMe',{'do': './install.py --gocode-completer --tern-completer --rust-completer  --clang-completer'}

" let g:ycm_filetype_blacklist = {'go':1}
let g:ycm_filetype_blacklist = {}
let g:ycm_auto_trigger = 1
let g:ycm_server_keep_logfiles = 1
let g:ycm_server_log_level = 'debug'
let g:ycm_min_num_of_chars_for_completion = 2
let g:ycm_key_list_select_completion = ['<c-j>','<Down>']
let g:ycm_key_list_previous_completion= ['<c-k>','<Up>']
let g:ycm_key_invoke_completion = '<c-z>'
" let g:ycm_python_binary_path = 'python2.7'
let g:ycm_confirm_extra_conf = 0
" nmap <M-g> :YcmCompleter GoToDefinitionElseDeclaration <C-R>=expand("<cword>")<CR><CR>  
Plug 'rdnetto/YCM-Generator', { 'branch': 'stable'}

Plug 'mdempsky/gocode', { 'rtp': 'vim', 'do': '~/.vim/plugged/gocode/vim/symlink.sh' }
"""""""""""""""""""""""""""""<code completion> end""""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""<code snippets> start""""""""""""""""""""""""""""""""""
" Track the engine.
Plug 'SirVer/ultisnips'

" Snippets are separated from the engine. Add this if you want them:
Plug 'honza/vim-snippets'

let g:UltiSnipsExpandTrigger="<tab>"

let g:UltiSnipsJumpForwardTrigger = "<c-l>"
let g:UltiSnipsJumpBackwardTrigger="<c-h>"
let g:UltiSnipsEditSplit="vertical"

"set runtimepath^=~/.config/nvim
"use absolute path 
" let g:UltiSnipsSnippetStorageDirectoryForUltiSnipsEdit = '~/.config/nvim/MySnippets'
" let g:UltiSnipsSnippetDirectories=["MySnippets","UltiSnips"]
let g:UltiSnipsSnippetStorageDirectoryForUltiSnipsEdit = '/Users/jialinwu/git/configuration/UltiSnips'
let g:UltiSnipsSnippetDirectories=["UltiSnips","/Users/jialinwu/git/configuration/UltiSnips"]
" let g:UltiSnipsSnippetDirectories=["UltiSnips","/Users/jialinwu/git/configuration/UltiSnips","/Users/jialinwu/.nvim/plugged/vim-snippets/snippets"]
let g:UltiSnipsSnippetsDir = "/Users/jialinwu/git/configuration/UltiSnips"
" let g:UltiSnipsSnippetsDir = '~/.config/nvim'
"""""""""""""""""""""""""""""<code snippets> end""""""""""""""""""""""""""""""""""""

"for js
Plug 'ternjs/tern_for_vim'
Plug 'pangloss/vim-javascript'
Plug 'posva/vim-vue'
Plug 'othree/html5.vim'

"rust
Plug 'rust-lang/rust.vim'

"auto format js html
" au BufNewFile,BufRead *.vue set filetype=html
autocmd BufWritePre *.{js,html,htm,vue} :normal migg=G`izz
" autocmd BufWritePre *.{js,html,htm} :normal ggVG=

" Markdown
autocmd BufNewFile,BufRead *.{md,mkd,mkdn,mark*}  nested setlocal filetype=markdown

"for django
au BufNewFile,BufRead *.hj set filetype=htmldjango

Plug 'godlygeek/tabular'
Plug 'plasticboy/vim-markdown'

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"neomake
Plug 'neomake/neomake'
autocmd! BufWritePost * Neomake
" autocmd BufWritePost * Neomake
let g:neomake_open_list=2
let g:neomake_go_enabled_makers = [ 'go', 'golangci_lint' ]
" let g:neomake_go_gometalinter_maker = {
"   \ 'args': [
"   \   '--tests',
"   \   '--enable-gc',
"   \   '--concurrency=3',
"   \   '--fast',
"   \   '-D', 'gocyclo',
"   \   '-D', 'gotype',
"   \   '-D', 'dupl',
"   \   '-D', 'golint',
"   \   '-D', 'ineffassign',
"   \   '-E', 'interfacer',
"   \   '-E', 'goconst',
"   \   '-E', 'aligncheck',
"   \   '-E', 'unconvert',
"   \   '-E', 'errcheck',
"   \   '-E', 'misspell',
"   \   '-E', 'unused',
"   \   '-D', 'vet',
"   \   '-D', 'vetshadow',
"   \   '%:p:h',
"   \ ],
"   \ 'append_file': 0,
"   \ 'errorformat':
"   \   '%E%f:%l:%c:%trror: %m,' .
"   \   '%W%f:%l:%c:%tarning: %m,' .
"   \   '%E%f:%l::%trror: %m,' .
"   \   '%W%f:%l::%tarning: %m'
"   \ }
" vue is treated as html file but tidy does'nt reconize vue syntax, so
" disabled tidy for vue safty
let g:neomake_html_enabled_makers = []
let g:neomake_css_enabled_makers = []

"tags lister method menu
Plug 'majutsushi/tagbar'
" let g:tagbar_width = 30
" let g:tagbar_left = 1
let g:tagbar_show_linenumbers = 2
let g:tagbar_type_markdown = {
        \ 'ctagstype' : 'markdown',
        \ 'kinds' : [
                \ 'h:headings',
        \ ],
    \ 'sort' : 0
\ }
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"shortcut for comment,e.g. gcc
Plug 'tomtom/tcomment_vim'
Plug 'tpope/vim-surround'
Plug 'tpope/vim-abolish'


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"json
" Plug 'elzr/vim-json'

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"power search tool
Plug 'mileszs/ack.vim'
if executable('ag')
  let g:ackprg = 'ag --vimgrep'
endif

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"show git diff when edicting
Plug 'airblade/vim-gitgutter'
let g:gitgutter_signs = 1
let g:gitgutter_highlight_lines = 0

"""""""""""""""""""""""""""""<golang tool chain> start""""""""""""""""""""""""""""""""""
Plug 'fatih/vim-go', { 'do': ':GoUpdateBinaries' }
let g:go_fmt_command = 'goimports'    "auto insert package
let g:go_def_mode='gopls'
let g:go_info_mode='gopls'
" let g:go_def_mode='godef'
"horizon
au FileType go nmap <Leader>h <Plug>(go-def-split) 
au FileType go nmap <Leader>v <Plug>(go-def-vertical)
au FileType go nmap <Leader>t <Plug>(go-def-tab)
"""""""""""""""""""""""""""""<golang tool chain> end""""""""""""""""""""""""""""""""""""

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" select multiple words and update them
Plug 'terryma/vim-multiple-cursors'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"scoll smoothly when hit c-b c-f
" Plug 'yonchu/accelerated-smooth-scroll'


" set dropdown to match solarized light
highlight Pmenu ctermfg=254 ctermbg=33
highlight PmenuSel ctermfg=235 ctermbg=39 cterm=bold
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"input stat switch between normal mode and insert mode automatically
Plug 'CodeFalling/fcitx-vim-osx'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" visual start search
Plug 'bronson/vim-visual-star-search'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'raimondi/delimitmate'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'ctrlpvim/ctrlp.vim'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'ryanoasis/vim-devicons'
let g:airline_powerline_fonts = 1

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'editorconfig/editorconfig-vim'
" let g:EditorConfig_exec_path ='~/.config/nvim/plugged/editorconfig-vim/plugin/editorconfig-core-py'
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" Plug 'justinmk/vim-sneak'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" Plug 'easymotion/vim-easymotion'
" Move to word
" map  w <Plug>(easymotion-bd-w)
" nmap w <Plug>(easymotion-overwin-w)
" s{char}{char} to move to {char}{char}
" nmap s <Plug>(easymotion-overwin-f2)
" Move to line
" map l <Plug>(easymotion-bd-jk)
" nmap l <Plug>(easymotion-overwin-line)

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

Plug 'cespare/vim-toml'
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'buoto/gotests-vim'
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" Plug 'hotoo/pangu.vim'
" autocmd BufWritePre *.markdown,*.md,*.text,*.txt,*.wiki,*.cnx call PanGuSpacing()
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'chase/vim-ansible-yaml'
" add yaml stuffs
au! BufNewFile,BufReadPost *.{yaml,yml} set filetype=yaml foldmethod=indent
autocmd FileType yaml setlocal ts=2 sts=2 sw=2 expandtab
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'altercation/vim-colors-solarized'

Plug 'sheerun/vim-polyglot'
Plug 'nvim-treesitter/nvim-treesitter', {'do': ':TSUpdate'}
"
" Plug 'shaunsingh/solarized.nvim'
" Plug 'maxmx03/solarized.nvim'

set background=light
" set background=dark
" let g:solarized_visibility = "high"
" let g:solarized_contrast = "high"
" let g:solarized_termcolors=256

" colorscheme molokai

" Important!!
if has('termguicolors')
	set termguicolors
endif

" "git clone --depth 1 https://github.com/lifepillar/vim-solarized8.git ~/.config/nvim/colors/
colorscheme solarized8


" let g:molokai_original = 1
" let g:rehash256 = 1
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
let g:vimrc_author='Jialin Wu' 
let g:vimrc_email='win27v@gmail.com' 
let g:vimrc_homepage='https://jialinwu.com'

Plug 'junegunn/fzf', { 'dir': '~/.fzf', 'do': './install --all' }
set rtp+=/Users/jialinwu/.fzf

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"powerline like
" Plug 'itchyny/lightline.vim'
" let g:lightline = {
"       \ 'colorscheme': 'solarized',
"       \ 'component_function': {
"       \   'filename': 'LightLineFilename'
"       \ },
"       \ }
function! LightLineFilename()
  " return expand('%:p')
  return expand('%')
endfunction


""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
" " To use solarized dark, set :AirlineTheme solarized and add the following to your .vimrc: let g:airline_solarized_bg='dark'
let g:airline_theme='solarized'
let g:airline_solarized_bg='light'
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

Plug 'Exafunction/codeium.vim', { 'branch': 'main' }
" let g:codeium_no_map_tab
let g:codeium_disable_bindings = 1

imap <script><silent><nowait><expr> <C-g> codeium#Accept()
imap <C-;>   <Cmd>call codeium#CycleCompletions(1)<CR>
imap <C-,>   <Cmd>call codeium#CycleCompletions(-1)<CR>
imap <C-x>   <Cmd>call codeium#Clear()<CR>

set statusline+=\{…\}%3{codeium#GetStatusString()}


Plug 'nvim-treesitter/nvim-treesitter', {'do': ':TSUpdate'}  " We recommend updating the parsers on update

Plug 'rhysd/git-messenger.vim' "It shows the history of commits under the cursor in popup window.

Plug 'tpope/vim-rsi'
call plug#end()

" ====================leader key=================
 "Set mapleader
let mapleader = "\<Space>"

"show file edicted recently 
" set runtimepath^=~/.config/nvim/bundle/ctrlp.vim
nnoremap <silent> <leader>r :CtrlPMRU<cr>
" map <silent> <leader>t :CtrlPMixed<cr>
" map <silent> <leader>t :CtrlP<cr>

"list file in current directory
nnoremap <silent> <leader>f :FZF<cr>

"Fast opening nerdtree 
nnoremap <silent> <leader>nt :NERDTreeToggle<cr>
nnoremap <silent> <leader>e :NERDTreeToggle<cr>

"exit without saving
map <silent> <leader>q :q!<cr>
map <silent> <leader>1q :qa!<cr>
"Save and exit
map <silent> <leader>x :x<cr>
"select all
map <silent> <leader>a ggvGV<cr>

"exit without saving
map <silent> <leader>q :q!<cr>
"exit all windows without saving
map <silent> <leader>qa :qa!<cr>

map <silent> <leader>b :NERDTreeFromBookmark 

noremap <Leader>d "+d
noremap <Leader>y "+y
noremap <Leader>1y "1y
noremap <Leader>2y "2y

noremap <Leader>p "+p
noremap <Leader>P "+P
noremap <Leader>1p "1p
noremap <Leader>1P "1P

noremap <Leader>2p "2p
noremap <Leader>2P "2P

"write zone
map <silent> <leader>wz :Goyo<cr>

"exit write zone
map <silent> <leader>ewz :Goyo!<cr>

"shwo absolute filename
nnoremap <silent> <leader>af :echo expand('%:p')<cr>

" quick into searching
nnoremap <silent> <leader>k :Ack 

"new tab 
nnoremap <silent> <leader>tn :tabnew<cr>
"close one tab 
nnoremap <silent> <leader>tc :tabclose<cr>

"switch window area
"for terminal of nvim
" :tnoremap <C-h> <C-\><C-n><C-w>h
" :tnoremap <C-j> <C-\><C-n><C-w>j
" :tnoremap <C-k> <C-\><C-n><C-w>k
" :tnoremap <C-l> <C-\><C-n><C-w>l
:tnoremap <C-\> <C-\><C-n>
"don't use this map while c-c is useful in terminal to operate shell cmd
" :tnoremap <C-c> <C-\><C-n>

:nnoremap <C-h> <C-w>h
:nnoremap <C-j> <C-w>j
:nnoremap <C-k> <C-w>k
:nnoremap <C-l> <C-w>l
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

nnoremap ; : 
nmap    w=  :resize +20<CR>
nmap    w-  :resize -20<CR>
nmap    w,  :vertical resize -20<CR>
nmap    w.  :vertical resize +20<CR>
"""""""""""""""""""""""""""""<shortcut defined> end""""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""<for cpp/c> start""""""""""""""""""""""""""""""""""
" switch between header/source with  hc for cpp
map <silent> <leader>hc :e %:p:s,.h$,.X123X,:s,.cpp$,.h,:s,.X123X$,.cpp,<cr>
"""""""""""""""""""""""""""""<for cpp/c> end""""""""""""""""""""""""""""""""""""

"edit tmp file
map <silent> <leader>tmp :e ~/tmpfile<cr>
map <silent> <leader>booknote :e /Users/Jialin/Dropbox/vimNote/book.md<cr>
map <silent> <leader>jsnote :e /Users/Jialin/Dropbox/vimNote/jsnote.md<cr>

"hightlight and search
vnoremap // y/<C-R>"<CR>

"make
map <silent> <leader>zm :make <cr>
map <silent> <leader>zl :lopen<cr>
map <silent> <leader>m :TagbarToggle<cr>

"tab to %
nnoremap <tab> %
vnoremap <tab> %
"keep hightlighting while doing indent 
vnoremap < <gv
vnoremap > >gv

"selec text from current cursor to the end of line
nnoremap Y v$

"delete text from current cursor to the end or begining of line in insert mode
map! <M-l>  <ESC>v$c
map! <M-h>  <ESC>v^c

"Keep search pattern at the center of the screen
nnoremap <silent> n nzz
nnoremap <silent> N Nzz
nnoremap <silent> * *zz
nnoremap <silent> # #zz
nnoremap <silent> g* g*zz

"Better comand-line editing
cnoremap <C-a> <Home>
cnoremap <C-e> <End>

"replace : with ; in normal mode
nnoremap ; :
"f{xx} forwrad`,` backward`<`
nnoremap ' ;
nnoremap " ,

"split window
nnoremap <neader>2w <C-w>v<C-w>l

"visul for html tab
nnoremap <leader>zv Vat

"use alt-w to save
nnoremap <M-w> :w<cr>
inoremap <M-w> <ESC>:w<cr>

"use jj to exit back to normal mode
inoremap lk <ESC>

"open dir of nerd-tree on current file
map <Leader>cd :NERDTree %:p:h<CR>

"Save file edicting
nnoremap <silent> <leader>w :w<cr>

"Fast editing of edictor configuration
map <silent> <leader>ee :e ~/.config/nvim/init.vim<cr>
autocmd! bufwritepost .nvim.rc source ~/.config/nvim/init.vim

hi Directory ctermfg=Blue
" hi Directory ctermfg=lightBlue
" hi Directory ctermfg=grey

"let g:python3_host_prog='/usr/local/bin/python3'
let g:python3_host_prog='/opt/homebrew/bin/python3'
" let g:python3_host_prog='/opt/homebrew/anaconda3/bin/python3'
"let g:python_host_prog='/usr/local/opt/python2/libexec/bin/python'
" let g:python3_host_prog='/usr/local/opt/python3/libexec/bin/python'
"let g:python_host_skip_check = 1
let g:tagbar_ctags_bin='/opt/homebrew/bin/ctags'

" lua require'nvim-treesitter.configs'.setup{highlight={enable=true}}  " At the bottom of your init.vim, keep all configs on one line
