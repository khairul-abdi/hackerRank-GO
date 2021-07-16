function name(params) {
  let count =0
    for (let i = 0; i < params.length; i++) {
      if (params[i] > params[i+1]) {
        let change = params[i]
        params[i] = params[i+1]
        params[i+1] = change
        count++
        i = -1
      } 
    }
  return count
}

console.log(name([4, 2, 3, 1]))