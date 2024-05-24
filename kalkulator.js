function kalkulator(a, b, c) {
  let d = 0;

  switch (c) {
    case (c = "+"):
      d = a + b;
      break;
    case (c = "-"):
      d = a - b;
      break;
    case (c = "/"):
      d = a / b;
      break;
    case (c = "*"):
      d = a * b;
      break;

    default:
      console.log("Error");
      return;
      break;
  }

  console.log("Hasil = " + d);
}
