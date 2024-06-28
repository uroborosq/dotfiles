-- AstroCore provides a central place to modify mappings, vim options, autocommands, and more!
-- Configuration documentation can be found with `:h astrocore`
-- NOTE: We highly recommend setting up the Lua Language Server (`:LspInstall lua_ls`)
--       as this provides autocomplete and documentation while editing

---@type LazySpec
return {
  "AstroNvim/astrocore",
  ---@type AstroCoreOpts
  opts = {

    -- vim options can be configured here
    options = {
      opt = {                   -- vim.opt.<key>
        relativenumber = false, -- sets vim.opt.relativenumber
        number = true,          -- sets vim.opt.number
        spell = false,          -- sets vim.opt.spell
        signcolumn = "auto",    -- sets vim.opt.signcolumn to auto
        wrap = false,
        tabstop = 4,
        shiftwidth = 4,
        softtabstop = 4,
        expandtab = true,
        cmdheight = 1, -- sets vim.opt.wrap
      },
    },

  },
}
