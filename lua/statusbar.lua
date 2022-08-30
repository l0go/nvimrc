local gl = require("galaxyline")
local gls = gl.section

local colors = require('rose-pine.plugins.galaxyline')
local condition = require('galaxyline.condition')
local fileinfo  = require('galaxyline.providers.fileinfo')

gls.left[1] = {
  ViMode = {
    provider = function()
      local alias = {
        n = "NORMAL",
        i = "INSERT",
        c = "COMMAND",
        v = "VISUAL",
        V = "VISUAL LINE",
        [""] = "VISUAL BLOCK",
      }
      return alias[vim.fn.mode()]
    end,
	separator = " • ",
	separator_highlight = { colors.yellow, colors.bg },
	highlight = { colors.magenta, colors.bg, "bold" },
  },
}

gls.left[2] = {
	FileName = {
		condition = condition.buffer_not_empty,
		highlight = {colors.white, colors.bg, 'bold'},
		provider  = function() return vim.fn.expand("%:t") end,
	}
}
