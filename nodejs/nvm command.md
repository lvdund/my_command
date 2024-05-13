In Node.js development, NVM (Node Version Manager) is a tool that allows you to easily manage multiple Node.js versions on your machine. It lets you switch between different versions of Node.js for different projects.

Here are some common NVM commands:

1. **Install NVM:**
   To install NVM, you can use the following command in your terminal:

   ```bash
   curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
   ```

   After the installation, restart your terminal or run the initialization script as instructed in the terminal.

2. **List available Node.js versions:**
   You can list all available Node.js versions that you can install using:

   ```bash
   nvm ls-remote
   ```

3. **Install a specific Node.js version:**
   To install a specific version of Node.js, you can use:

   ```bash
   nvm install <version>
   ```

   For example:

   ```bash
   nvm install 14.17.0
   ```

4. **Set default Node.js version:**
   To set a default Node.js version to be used in new shells, you can use:

   ```bash
   nvm alias default <version>
   ```

   For example:

   ```bash
   nvm alias default 14.17.0
   ```

5. **Use a specific Node.js version:**
   To use a specific version of Node.js in the current shell session, you can use:

   ```bash
   nvm use <version>
   ```

   For example:

   ```bash
   nvm use 14.17.0
   ```

6. **List installed Node.js versions:**
   You can list all locally installed Node.js versions with:

   ```bash
   nvm ls
   ```

7. **Uninstall a Node.js version:**
   To uninstall a specific version of Node.js, you can use:

   ```bash
   nvm uninstall <version>
   ```

   For example:

   ```bash
   nvm uninstall 14.17.0
   ```

These commands should help you manage your Node.js versions efficiently using NVM.
