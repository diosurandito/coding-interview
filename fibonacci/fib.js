const fib = (n) => {
    if (n > 1) {
        return fib(n - 1) + fib(n - 2);
    } else if (n === 0){
        return 0;
    } else if (n === 1){
        return 1;
    } else {
        return "invalid";
    }
};

function printFib(num){
    fs = [];
    for (let i = 0; i < num; i++) {
        fs.push(fib(i));
    }
    return fs;
}

console.log(printFib(8));