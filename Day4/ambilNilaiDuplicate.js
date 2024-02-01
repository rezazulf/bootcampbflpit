var input = [10, 9, 10, 9, 8, 5, 5, 9, 6, 6, 'a', 'c', 'd', 'e', 'c', 'd', 'f', 'c', 'c'];

function count(arr) {
    var result = arr.map(function (element, index) {
        var count = arr.filter(function (el) {
            return el === element;
        }).length;

        return { element: element, count: count };
    });

    var maxCount = Math.max(...result.map(item => item.count));
    var mostFrequent = result.find(item => item.count === maxCount);

    return [mostFrequent.element, mostFrequent.count];
}

var output = count(input);
console.log(output);
