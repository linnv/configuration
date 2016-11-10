syntax on  			"must put ahead or enable syntax will conflict with ultisnips
set ruler "display the current row and column on the right_bottom corner 
set relativenumber
set nu

" set foldmethod=syntax "fold based on indent
set foldnestmax=10      "deepest fold is 10 levels
set nofoldenable        "dont fold by default
set foldlevel=1         "this is just what i use

set fillchars+=vert:│
" set fillchars+=vert:\ 

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

"remember last edicting position
au BufReadPost * if line("'\"") > 0|if line("'\"") <= line("$")|exe("norm '\"")|else|exe "norm $"|endif|endif

filetype off                  " required: ultisnipes
" ====================plugin manager=================
call plug#begin('~/.config/nvim/plugged')
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'scrooloose/nerdtree'
"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""<code completion> start""""""""""""""""""""""""""""""""""
Plug 'Valloric/YouCompleteMe',{'do': './install.py --tern-completer  --clang-completer'}

let g:ycm_auto_trigger = 1
let g:ycm_min_num_of_chars_for_completion = 2
let g:ycm_key_list_select_completion = ['<c-j>','<Down>']
let g:ycm_key_list_previous_completion= ['<c-k>','<Up>']
let g:ycm_key_invoke_completion = '<c-z>'

let g:ycm_confirm_extra_conf = 0
let g:ycm_global_ycm_extra_conf ='~/.config/nvim/plugged/YouCompleteMe/third_party/ycmd/cpp/ycm/.ycm_extra_conf.py' 
nmap <M-g> :YcmCompleter GoToDefinitionElseDeclaration <C-R>=expand("<cword>")<CR><CR>  
Plug 'rdnetto/YCM-Generator'
"""""""""""""""""""""""""""""<code completion> end""""""""""""""""""""""""""""""""""""

"""""""""""""""""""""""""""""<code snippets> start""""""""""""""""""""""""""""""""""
Plug 'SirVer/ultisnips' | Plug 'honza/vim-snippets'
let g:UltiSnipsExpandTrigger="<tab>"
let g:UltiSnipsJumpForwardTrigger = "<c-l>"
let g:UltiSnipsJumpBackwardTrigger="<c-h>"
let g:UltiSnipsEditSplit="vertical"
"use absolute path 
let g:UltiSnipsSnippetsDir = '~/.config/nvim/UltiSnips'
"""""""""""""""""""""""""""""<code snippets> end""""""""""""""""""""""""""""""""""""

"for js
 Plug 'ternjs/tern_for_vim'
 Plug 'pangloss/vim-javascript'

" Markdown
autocmd BufNewFile,BufRead *.{md,mkd,mkdn,mark*}  nested setlocal filetype=markdown
Plug 'godlygeek/tabular'
Plug 'plasticboy/vim-markdown'

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"neomake
Plug 'neomake/neomake'
autocmd! BufWritePost * Neomake
" let g:neomake_open_list=2
" let g:neomake_go_enabled_makers = ['go', 'govet']

"tags lister method menu
Plug 'majutsushi/tagbar'
" let g:tagbar_width = 30
" let g:tagbar_left = 1
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

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"power search tool
Plug 'mileszs/ack.vim'

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"show git diff when edicting
Plug 'airblade/vim-gitgutter'
let g:gitgutter_signs = 1
let g:gitgutter_highlight_lines = 0

"""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"color symble that pairs each other
Plug 'luochen1990/rainbow'
let g:rainbow_active= 1

"""""""""""""""""""""""""""""<golang tool chain> start""""""""""""""""""""""""""""""""""
Plug 'fatih/vim-go'
let g:go_fmt_command = 'goimports'    "auto insert package
"horizon
au FileType go nmap <Leader>h <Plug>(go-def-split) 
au FileType go nmap <Leader>v <Plug>(go-def-vertical)
au FileType go nmap <Leader>t <Plug>(go-def-tab)
"""""""""""""""""""""""""""""<golang tool chain> end""""""""""""""""""""""""""""""""""""

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
set background=light
colorscheme solarized

" colorscheme molokai
" let g:molokai_original = 1
" let g:rehash256 = 1

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
" select multiple words and update them
Plug 'terryma/vim-multiple-cursors'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"scoll smoothly when hit c-b c-f
" Plug 'yonchu/accelerated-smooth-scroll'

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
"powerline like
Plug 'itchyny/lightline.vim'
let g:lightline = {
      \ 'colorscheme': 'solarized',
      \ 'component_function': {
      \   'filename': 'LightLineFilename'
      \ },
      \ }
function! LightLineFilename()
  " return expand('%:p')
  return expand('%')
endfunction
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

""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
Plug 'buoto/gotests-vim'
""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""""
let g:vimrc_author='Jialin Wu' 
let g:vimrc_email='win27v@gmail.com' 
let g:vimrc_homepage='https://jialinwu.com'

call plug#end()
" call vundle#end()            " required
filetype plugin indent on    " required

" ====================leader key=================
 "Set mapleader
let mapleader = "\<Space>"

"show file edicted recently 
set runtimepath^=~/.config/nvim/bundle/ctrlp.vim
nnoremap <silent> <leader>r :CtrlPMRU<cr>
" map <silent> <leader>t :CtrlPMixed<cr>
" map <silent> <leader>t :CtrlP<cr>

"list file in current directory
Plug 'junegunn/fzf', { 'dir': '~/.fzf', 'do': './install --all' }
set rtp+=/usr/local/opt/fzf
nnoremap <silent> <leader>f :FZF<cr>

"Fast opening nerdtree 
nnoremap <silent> <leader>nt :NERDTree<cr>

"Save and exit
map <silent> <leader>x :x<cr>
"select all
map <silent> <leader>a ggvGV<cr>

"exit without saving
map <silent> <leader>q :q!<cr>

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

"shwo absolute filename
nnoremap <silent> <leader>af :echo expand('%:p')<cr>

" quick into searching
nnoremap <silent> <leader>k :Ack 

"new tab 
nnoremap <silent> <leader>tn :tabnew<cr>

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
map <silent> <leader>zm :make <cr>
map <silent> <leader>zl :lopen<cr>
map <silent> <leader>m :TagbarToggle<cr>

"tab to %
nnoremap <tab> %
vnoremap <tab> %

"replace : with ; in normal mode
nnoremap ; :
"f{xx} forwrad`,` backward`<`
nnoremap , ;
nnoremap < ,

"split window
nnoremap <leader>2w <C-w>v<C-w>l

"visul for html tab
nnoremap <leader>zv Vat

"use alt-w to save
nnoremap <M-w> :w<cr>
inoremap <M-w> <ESC>:w<cr>

"use jj to exit back to normal mode
inoremap kl <ESC>

"Save file edicting
nnoremap <silent> <leader>w :w<cr>
"""""""""""""""""""""""""""""<for nvim configuration> start""""""""""""""""""""""""""""""""""
"Fast editing of edictor configuration
map <silent> <leader>ee :e ~/.config/nvim/init.vim<cr>
autocmd! bufwritepost init.vim source ~/.config/nvim/init.vim

hi Directory ctermfg=Blue

let g:python_host_prog='/usr/bin/python2.7'
let g:python3_host_prog='/usr/local/bin/python3'
let g:python_host_skip_check = 1
"""""""""""""""""""""""""""""<for nvim configuration> end""""""""""""""""""""""""""""""""""""
