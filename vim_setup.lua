local client = vim.lsp.start_client {
    name = "test_lsp",
    cmd = { "" }, //replace with the lsp server file
    on_attach = {},
}

if not client then
    vim.notify "hey you did not run the client"
    return
end

vim.api.nvim_create_autocmd("FileType", {
    pattern = "markdown",
    callback = function()
        vim.lsp.buf_attach_client(0, client)
    end,
})
