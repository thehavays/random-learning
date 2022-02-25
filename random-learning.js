const {Command} = require('commander');
const program = new Command();

program
    .name('random-learning')
    .description('Random terms and titles for learning.')
    .version('0.0.1', '-v, --vers', 'Output the current version');

program.command('write')
    .description('It\'s a example command for testing.')
    .argument('<string>', 'string to write')
    .option('-l, --isLength <boolean>', 'write parameter length', 'false')
    .action((str, options) => {
        const isLength = options.isLength === 'true'
        const resultString = isLength ? str + "aaa" : str
        console.log(resultString)
    });

program.parse();