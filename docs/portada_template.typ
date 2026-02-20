
#let template(doc) = [
  #set page(
    paper: "a4",
    //margin: 2cm,
    header: none,
    // numbering: "1",
    number-align: right,
  )
  
  #show link: it => [
    #set text(blue)
    //underline(blue, thickness: 0.05em)
    #it
  ]
  
  #set quote(block: true)
  #show quote: set pad(x: 2em)
  
  #set bibliography(title: none)
  
  #doc
]

#let portada(clase, actividad, nombre, matricula, carrera, fecha, espacio) = {

  set page(
  paper: "a4",
  margin: 2cm,
  header: none,
  //numbering: "1",
  // number-align: right,
  )


  set text(font: "Lato", 12pt)
  set quote(block: true)
  show quote: set pad(x: 2em)

 page()[
  
  // Logos de UANL + FIME
  #grid(
    columns: (1fr, 1fr),
    column-gutter: 0pt,
    align(left)[
      #image("img/uanl.png", width: 4cm)
    ],
    align(right)[
      #image("img/fime.png", width: 4cm)
    ]
  )

  
  #v(0.9cm)
  
  
  #align(center)[
    
    #text(size: 24pt, weight: "regular")[Universidad Autónoma de Nuevo León
] \
    #v(0.5cm)
    #text(size: 23pt, weight: "regular")[Facultad de Ingeniería Mecánica y Eléctrica ] \

    
    #v(2cm)
    
    
    #text(size: 18pt)[#text(weight: "regular")[#clase]] \
    #v(0.5cm)
    #text(size: 18pt)[#text(weight: "regular")[#actividad]] \
    #v(2cm)
  ]

  #table(
  columns: (1fr, 1fr, 1fr),   
  stroke: black,              
  align: center,              
  // Si es enquipo pon aqui datos de equipo
    [*Nombre*], [*Matricula*], [*Carrera*],
    [#nombre], [#matricula], [#carrera],
    [Integrante 2], [Matricula 2], [x],
    [Integrante 3], [Matricula 3], [x],
    [Integrante 4], [Matricula 4], [x],
    [Integrante 5], [Matricula 5], [x],
  )

  // manualmente cambiar espacio (deberia hacer esto con footer)
  #v(espacio) 

  // footer
  #align(center)[
    Ciudad Universitaria, San Nicolás de los Garza, Nuevo León 
    
    #fecha.display()
  ]
  #pagebreak()
] 
}


