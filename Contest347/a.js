function main (input){
    console.log(input);
    // input = input.split('\n');
    
    // const [n, a, b] = input.split(' ').map(Number);
    // console.log(input);
}

main(require('fs').readFileSync('./input.txt'), 'utf8');