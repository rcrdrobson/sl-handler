module.exports.helloWorld = (req, res) => {
    res.send("hello world")
}

module.exports.fib = (req, res) => {
    res.send(""+fib(req.query.value))
}

function fib(n){
    if(n < 2) return n;
    return fib(n - 1) + fib(n - 2);
}

/*

module.exports.helloWorld = (req, res) => {\r\n    res.send(\"hello world\")\r\n}\r\n\r\nmodule.exports.fib = (req, res) => {\r\n    res.send(fib(req.query.value))\r\n}\r\n\r\nfunction fib(n){\r\n    if(n < 2) return n;\r\n    return fib(n - 1) + fib(n - 2);\r\n}

*/