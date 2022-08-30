colorscheme rose-pine

set autoindent
set number
set ruler
set hlsearch
set history=100
set omnifunc=syntaxcomplete#Complete
let maplocalleader = ','
set tabstop=4
set softtabstop=0 noexpandtab
set shiftwidth=4

syntax on
filetype plugin indent on
filetype plugin on

nnoremap <silent> . :Files<CR>

set runtimepath+=~/.config/nvim/
lua require('statusbar')
