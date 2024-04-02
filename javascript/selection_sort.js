const findSmallest = (arr) => {
  let smallest = arr[0];
  let smallest_index = 0;
  for (let i = 1; i <= arr.length; i++) {
    if (arr[i] < smallest) {
      smallest = arr[i];
      smallest_index = i;
    }
  }
  return smallest_index;
};

const selectionSort = (arr) => {
  newArr = [];
  const len = arr.length;
  for (let i = 1; i <= len; i++) {
    smallest = findSmallest(arr);
    newArr.push(arr.splice(smallest, 1)[0]);
  }
  return newArr;
};

console.log(selectionSort([5, 3, 6, 2, 10]));
