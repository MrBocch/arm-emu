
s = ARGV[0]

puts s.chars.map{ |c|
  log = ""
  log += "// output #{c}\n"
  log += "mov r0, ##{c.ord}\nsys #3\n"
}
