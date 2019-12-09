let {exec} = require('child_process');
let fs = require('fs');
getInput(process.argv[2], process.argv[3]);
function getInput(day, year){
    if (!day) return console.log('[ERR] Day is undefined!');
    if (!year) year = 2019;
    let cookie = process.env.aoccookie;
    let cmd = `curl https://adventofcode.com/${year}/day/${day}/input --cookie session=${cookie}`;
    return exec(cmd, (err, stdout, info) => {
        if (err) return console.log(`[ERR] ${err}`);
        let errmsg = "Please don't repeatedly request this endpoint before it unlocks! The calendar countdown is synchronized with the server time; the link will be enabled on the calendar the instant this puzzle becomes available.\n";
        if (stdout == errmsg) return console.log('[ERR] N/A yet!');
        if (stdout == '404 Not Found\n') return console.log('[ERR] Error 404!');
        let filename = 'input';
        return fs.writeFile(`${process.cwd()}\\${filename}`, stdout.substring(0, stdout.length-1), () => console.log(`[OK!] ${day}/${year}\n[OK!] ${process.cwd()}\\${filename} (${parseInt(stdout.length-1)})`));
    });
}