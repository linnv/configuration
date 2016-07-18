syntax on  			"must put ahead or enable syntax will conflict with ultisnips
set ruler "display the current row and column on the right_bottom corner 
set relativenumber
set nu
set foldmethod=syntax "fold based on indent
set foldnestmax=10      "deepest fold is 10 levels
set nofoldenable        "dont fold by default
set foldlevel=1         "this is just what i use

set fillchars=stl:-,stlnc:-,vert:\ 

" ignorecase and smartcase together make Vim deal with case-sensitive search intelligently. If you search for an all-lowercase string your search will be case-insensitive, but if one or more characters is uppercase the search will be case-sensitive. Most of the time this does what you want.
set ignorecase
set smartcase

set showmatch
set showmode
set showcmd
set gdefault

set colorcolumn=85
set wrap
 set mouse=a

set timeoutlen=1000 ttimeoutlen=0

set modelines=0   "don't execute command that comment in a file"
set undofile  "this allow you to undo change even when you reopen a file
set laststatus=2   " Always show the statusline

" set cursorline  
set cursorcolumn
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
set t_Co=256
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"remember last edicting position
au BufReadPost * if line("'\"") > 0|if line("'\"") <= line("$")|exe("norm '\"")|else|exe "norm $"|endif|endif
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

filetype off                  " required: ultisnipes
" ====================plugin manager=================
set rtp+=~/.config/nvim/bundle/Vundle.vim
call vundle#begin('~/.config/nvim/bundle')
Plugin 'gmarik/Vundle.vim'
Plugin 'scrooloose/nerdtree'
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"""""""""""""""""""""""""""""<code completion> start""""""""""""""""""""""""""""""""""
Plugin 'Valloric/YouCompleteMe'
let g:ycm_key_list_select_completion = ['<c-j>','<Down>']
let g:ycm_key_list_previous_completion= ['<c-k>','<Up>']
let g:ycm_key_invoke_completion = '<c-z>'

let g:ycm_confirm_extra_conf = 0
let g:ycm_collect_identifiers_from_tags_files = 1
set tags+=./.tags
"""""""""""""""""""""""""""""<code completion> end""""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""<code snippets> start""""""""""""""""""""""""""""""""""
Plugin 'SirVer/ultisnips'
"Plugin 'honza/vim-snippets'
let g:UltiSnipsExpandTrigger="<tab>"
let g:UltiSnipsJumpForwardTrigger = "<c-l>"
let g:UltiSnipsJumpBackwardTrigger="<c-h>"
let g:UltiSnipsEditSplit="vertical"
""let g:UltiSnipsSnippetsDir = '~/.config/nvim/UltiSnips'
"""""""""""""""""""""""""""""<code snippets> end""""""""""""""""""""""""""""""""""""

" Markdown
autocmd BufNewFile,BufRead *.{md,mkd,mkdn,mark*}  nested setlocal filetype=markdown
Plugin 'godlygeek/tabular'
Plugin 'plasticboy/vim-markdown'

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"neomake
Plugin 'neomake/neomake'
autocmd! BufWritePost * Neomake
" let g:neomake_go_enabled_makers = ['go', 'govet']

"""""""""""""""""""""""""""""<syntax checker> start""""""""""""""""""""""""""""""""""
"use neomake instead
" Plugin 'scrooloose/syntastic'
" let g:syntastic_error_symbol = '✗'      "set error or warning signs
" let g:syntastic_warning_symbol = '⚠'
" let g:syntastic_enable_signs=1
" let g:syntastic_always_populate_loc_list = 1
" "
" let g:syntastic_enable_highlighting = 0
"""""""""""""""""""""""""""""<syntax checker> end""""""""""""""""""""""""""""""""""""

"tags lister method menu
Plugin 'majutsushi/tagbar'
nmap <c-m> :TagbarToggle<CR>


"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"short comment,e.g. gcc
Plugin 'tomtom/tcomment_vim'
Plugin 'tpope/vim-surround'

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"power search tool
Plugin 'mileszs/ack.vim'

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"show git diff when edicting
Plugin 'airblade/vim-gitgutter'
let g:gitgutter_signs = 1
let g:gitgutter_highlight_lines = 0

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"color symble that pairs
Plugin 'luochen1990/rainbow'
let g:rainbow_active= 1

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plugin 'rdnetto/YCM-Generator'

"""""""""""""""""""""""""""""<golang tool chain> start""""""""""""""""""""""""""""""""""
Plugin 'fatih/vim-go'
let g:go_fmt_command = 'goimports'    "auto insert package
au FileType go nmap <Leader>s <Plug>(go-def-split)
au FileType go nmap <Leader>v <Plug>(go-def-vertical)
au FileType go nmap <Leader>t <Plug>(go-def-tab)
"""""""""""""""""""""""""""""<golang tool chain> end""""""""""""""""""""""""""""""""""""

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"theme
colorscheme molokai
let g:molokai_original = 1
let g:rehash256 = 1

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" select multiple words and update them
Plugin 'terryma/vim-multiple-cursors'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"scoll smoothly when hit c-b c-f
" Plugin 'yonchu/accelerated-smooth-scroll'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"write in simple background
Plugin 'junegunn/goyo.vim'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"also say powerline
Plugin 'bling/vim-airline'


""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"input stat switch between normal mode and insert mode automatically
Plugin 'CodeFalling/fcitx-vim-osx'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" visual start search
Plugin 'bronson/vim-visual-star-search'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
let g:vimrc_author='Jialin Wu' 
let g:vimrc_email='win27v@gmail.com' 
let g:vimrc_homepage='http://linnv.gitcafe.io' 

call vundle#end()            " required
filetype plugin indent on    " required

" ====================leader key=================
 "Set mapleader
"let mapleader = ","
let mapleader = "\<Space>"

"""""""""""""""""""""""""""""<show file edicted recently> start""""""""""""""""""""""""""""""""""
" set runtimepath^=~/.vim/bundle/ctrlp.vim
set runtimepath^=~/.config/nvim/bundle/ctrlp.vim
map <silent> <leader>r :CtrlPMRU<cr>
" map <silent> <leader>t :CtrlPMixed<cr>
" map <silent> <leader>t :CtrlP<cr>
"""""""""""""""""""""""""""""<show file edicted recently> end""""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""<list file in current directory> start""""""""""""""""""""""""""""""""""
set rtp+=/usr/local/opt/fzf
map <silent> <leader>f :FZF<cr>
"""""""""""""""""""""""""""""<list file in current directory> end""""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""<shortcut defined> start""""""""""""""""""""""""""""""""""
"Fast opening nerdtree 
map <silent> <leader>nt :NERDTree<cr>

"Save file edicting
map <silent> <leader>w :w<cr>

"Save and exit
map <silent> <leader>x :x<cr>
"select all
map <silent> <leader>a ggvGV<cr>

"exit without saving
map <silent> <leader>q :q!<cr>
vmap <Leader>y "+y
vmap <Leader>1y "1y
vmap <Leader>2y "2y

vmap <Leader>d "+d
nmap <Leader>p "+p
nmap <Leader>P "+P
vmap <Leader>p "+p
vmap <Leader>P "+P

nmap <Leader>1p "1p
nmap <Leader>1P "1P
vmap <Leader>1p "1p
vmap <Leader>1P "1P

nmap <Leader>2p "2p
nmap <Leader>2P "2P
vmap <Leader>2p "2p
vmap <Leader>2P "2P

nnoremap <leader>jd :YcmCompleter GoToDefinitionElseDeclaration<CR>
nnoremap <F5> :YcmForceCompileAndDiagnostics<CR>
"write zone
map <silent> <leader>wz :Goyo<cr>
"exit write zone
map <silent> <leader>ewz :Goyo!<cr>

"switch window area
map <C-j> <C-W>j
map <C-k> <C-W>k
map <C-h> <C-W>h
map <C-l> <C-W>l
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

nnoremap ; : 
vnoremap // y/<C-R>"<CR>
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
map <silent> <leader>m :make <cr>
"""""""""""""""""""""""""""""<for nvim configuration> start""""""""""""""""""""""""""""""""""
"Fast editing of edictor configuration
map <silent> <leader>ee :e ~/.config/nvim/init.vim<cr>
autocmd! bufwritepost init.vim source ~/.config/nvim/init.vim

hi Directory ctermfg=lightBlue

let g:python_host_prog='/usr/bin/python2.7'
let g:python_host_skip_check = 1
    if !has('nvim')
        set ttymouse=xterm2
    endif
    if has('nvim')
     	nmap <BS> <C-W>h
    endif
"""""""""""""""""""""""""""""<for nvim configuration> end""""""""""""""""""""""""""""""""""""
