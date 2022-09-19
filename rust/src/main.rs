

fn next(before: String) -> String {
  let mut current_ch = before.chars().nth(0).unwrap();
  let mut count = 1;
  let mut s = "".to_owned();
  for ch in before.chars().skip(1) {
    if ch == current_ch {
      count += 1;
    } else {
      let next_str = format!("{}{}", current_ch, count);
      s.push_str(&next_str);
      current_ch = ch;
      count = 1;
    }
  }
  let next_str = format!("{}{}", current_ch, count);
  s.push_str(&next_str);

  s
}

fn ant() -> String {
  let mut s = "1".to_owned();
  for _ in 0..80 {
    s = next(s);
  }
  s
}

fn main() {
  let s = ant();
  println!("{}", s.chars().nth(0).unwrap());
}
