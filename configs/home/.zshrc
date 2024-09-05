# Enable Powerlevel10k instant prompt. Should stay close to the top of ~/.zshrc.
# Initialization code that may require console input (password prompts, [y/n]
# confirmations, etc.) must go above this block; everything else may go below.
if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
  source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
fi


# The following lines were added by compinstall
zstyle :compinstall filename "$HOME/.zshrc"

autoload -Uz compinit
compinit
# End of lines added by compinstall

# Lines configured by zsh-newuser-install
HISTFILE=~/.zsh_history
HISTSIZE=1000
SAVEHIST=10000
bindkey -e

source /usr/share/zsh/share/antigen.zsh
# End of lines configured by zsh-newuser-install


# Antigen configuration part

antigen bundle git
antigen bundle heroku
antigen bundle pip
antigen bundle lein
antigen bundle command-not-found

antigen bundle zsh-users/zsh-history-substring-search
antigen bundle zdharma-continuum/fast-syntax-highlighting
antigen bundle zsh-users/zsh-autosuggestions

antigen theme romkatv/powerlevel10k


antigen apply

# Keybinds

bindkey '^[[A' history-substring-search-up
bindkey '^[[B' history-substring-search-down
bindkey '^[[H' beginning-of-line
bindkey '^[[F' end-of-line

typeset -g -A key

key[Home]="${terminfo[khome]}"
key[End]="${terminfo[kend]}"
key[Insert]="${terminfo[kich1]}"
key[Backspace]="${terminfo[kbs]}"
key[Delete]="${terminfo[kdch1]}"
key[Up]="${terminfo[kcuu1]}"
key[Down]="${terminfo[kcud1]}"
key[Left]="${terminfo[kcub1]}"
key[Right]="${terminfo[kcuf1]}"
key[PageUp]="${terminfo[kpp]}"
key[PageDown]="${terminfo[knp]}"
key[Shift-Tab]="${terminfo[kcbt]}"

[[ -n "${key[Home]}"      ]] && bindkey -- "${key[Home]}"       beginning-of-line
[[ -n "${key[End]}"       ]] && bindkey -- "${key[End]}"        end-of-line
[[ -n "${key[Insert]}"    ]] && bindkey -- "${key[Insert]}"     overwrite-mode
[[ -n "${key[Backspace]}" ]] && bindkey -- "${key[Backspace]}"  backward-delete-char
[[ -n "${key[Delete]}"    ]] && bindkey -- "${key[Delete]}"     delete-char
[[ -n "${key[Up]}"        ]] && bindkey -- "${key[Up]}"         up-line-or-history
[[ -n "${key[Down]}"      ]] && bindkey -- "${key[Down]}"       down-line-or-history
[[ -n "${key[Left]}"      ]] && bindkey -- "${key[Left]}"       backward-char
[[ -n "${key[Right]}"     ]] && bindkey -- "${key[Right]}"      forward-char
[[ -n "${key[PageUp]}"    ]] && bindkey -- "${key[PageUp]}"     beginning-of-buffer-or-history
[[ -n "${key[PageDown]}"  ]] && bindkey -- "${key[PageDown]}"   end-of-buffer-or-history
[[ -n "${key[Shift-Tab]}" ]] && bindkey -- "${key[Shift-Tab]}"  reverse-menu-complete

bindkey ';5C' forward-word
bindkey ';5D' backward-word

# Aliases definitions part

alias ok=exit
alias nf=fastfetch
alias vim='nvim'
alias svim='sudo nvim'
alias pg='pwgen -cn1sy 16'
alias ls="ls --color"
alias l="lsd --date '+%d.%m.%Y %H:%M' -lah"
alias св='cd'
alias rmunsed='sudo pacman -Rnucs $(pacman -Qtdq)'
alias grep='grep --color=auto'
alias ssh='TERM=xterm-256color ssh'
alias sssh='TERM=xterm-256color ssh -oStrictHostKeyChecking=no -oUserKnownHostsFile=/dev/null'
# Path

export PATH="$HOME/.local/bin:$PATH"
export PATH="$HOME/Рабочий стол/Linux/bin/:$PATH"
export PATH="$HOME/Рабочий стол/Linux/scripts/:$PATH"
export PATH="$HOME/.local/share/go/bin/:$PATH"

ulimit -n 10000

if [ -n "$NVIM" ]; then
    NVIM_LISTEN_ADDRESS="$NVIM"
    alias nvim=nvr -cc split --remote-wait +'set bufhidden=wipe'
fi

if [ -n "$NVIM" ]; then
    export NVIM_LISTEN_ADDRESS="$NVIM"
    export VISUAL="nvr -cc split --remote-wait +'set bufhidden=wipe'"
    export EDITOR="nvr -cc split --remote-wait +'set bufhidden=wipe'"
else
    export VISUAL="nvim"
    export EDITOR="nvim"
fi
# To customize prompt, run `p10k configure` or edit ~/.p10k.zsh.
[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh
