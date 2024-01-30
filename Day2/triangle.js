let segitiga = "";
for (i = 5; i > 0; i--) {
  for (j = 0; j <= 5; j++) {
    if (j >= i) {
      segitiga += "*";
    } else {
      segitiga += " ";
    }
  }
  segitiga += "\n";
}
console.log(segitiga);


function segitiga2(sisi) {
  let result = '';
  for (let i = sisi; i > 0; i--) {
      for (let j = 1; j <= sisi; j++) {
          if (i <= j && j % 2==0) {
              result += '# ';
          }
          else if (i <= j && j % 2==1) {
                  result += '* ';
          } else {
              result += '  '
          }
      }
      result += '\n';
  }
  return result;
}
console.log(segitiga2(5));


for (let i = 1; i <= 5; i++) {
  let segitigas = '';
  for (let j = 1; j <= 5 - i; j++) {
    segitigas += ' ';
  }
  for (let k = 1; k <= 2 * i - 1; k++) {
    segitigas += '*';
  }
  console.log(segitigas);
}