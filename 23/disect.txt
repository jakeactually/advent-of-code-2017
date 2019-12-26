set b 93                
set c b                 
jnz a 2                 
jnz 1 5
mul b 100               
sub b -100000           
set c b                 
sub c -17000            
set f 1                 ; for i in 0..1000; b = 109300 + 17 * i; f = 1
set d 2                     ; d = 2
set e 2                     ; do ; e = 2
set g d                         ; do
mul g e                             
sub g b                             
jnz g 2                             ; if d * e == b
set f 0                                 ; f = 0
sub e -1                            ; e++
set g e                             
sub g b                             
jnz g -8                        ; while e != b
sub d -1                    ; d++
set g d                     
sub g b                     
jnz g -13                   ; while d != b
jnz f 2                     ; if f == 0
sub h -1                        ; h++
set g b                     
sub g c                     
jnz g 2                     
jnz 1 3                         
sub b -17                   
jnz 1 -23               
