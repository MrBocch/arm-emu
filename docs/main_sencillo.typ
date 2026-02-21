#import "portada_template.typ": *

#let fecha = datetime(
  year:  2026,
  month: 1,
  day:   1,
)

#portada(
    "Proyecto Integrador II",
    "Equipo #1",
    "Jorge Alberto Morales Reyes",
    "1895340",
    "ITS",
    fecha,
    5.5cm
)

#set document(
  title: [Proyecto Integrador II],
  author: "Jorge Alberto Morales Reyes/jorge.sh"
)

#set par(justify: true)

#set heading(numbering: "1.")
//#set text(font: "New Computer Modern")

#outline(
    title: [Indice],
)
#set page(
  numbering: "1",
  number-align: right,
)

#let header_bcolor = rgb(184, 209, 191)

#pagebreak()

= Definicion del Proyecto
== Nombre del Proyecto

[NOMBRE DE PROYECTO]

Razones porque? 

== Descripcion del Proyecto

El proyecto se concibe como un sistema compuesto por cuatro elementos principales: una *máquina virtual* basada en una arquitectura ARM32 simplificada, un *assembler* encargado de traducir código ensamblador a una representación binaria, una *interfaz gráfica* que permite la visualización y control del sistema, y una documentación técnica presentada en forma de manual. Estos componentes trabajan de manera conjunta para permitir la ejecución y el análisis de programas a bajo nivel en un entorno controlado, visual e interactivo.

== Problematica 

El aprendizaje de la arquitectura de computadoras, y en particular de
arquitecturas ampliamente utilizadas como ARM32, presenta diversas
dificultades en el ámbito educativo. Una de las principales problemáticas es la
limitada disponibilidad de hardware adecuado para la enseñanza. Aunque
existen dispositivos basados en ARM, como microcontroladores o placas de
desarrollo, su acceso no siempre es inmediato, su configuración puede resultar
compleja y, en muchos casos, su uso requiere conocimientos previos que van
más allá de los objetivos de cursos introductorios.

Adicionalmente, el estudio directo sobre hardware real suele desplazar el
enfoque del aprendizaje hacia detalles específicos de implementación, como
configuraciones del sistema, periféricos o particularidades del entorno de
ejecución. Esto puede dificultar que los estudiantes comprendan la visión
general de la arquitectura de un CPU, es decir, cómo interactúan los registros,
la memoria, el ciclo de ejecución y el conjunto de instrucciones,
independientemente de una plataforma física concreta.

Desde la perspectiva de un tutor o docente, estas dificultades representan un
reto importante en el proceso de enseñanza de la arquitectura de
computadoras. En consecuencia, el tutor se enfrenta a la necesidad de contar
con una herramienta que permita aislar los conceptos esenciales de la
arquitectura de procesadores, facilite la experimentación controlada y sirva
como apoyo didáctico tanto en el aula como en el estudio autónomo del
estudiante. Estos son los problemas educativos que el presente proyecto busca
abordar.

== Propuesta

Para atender las necesidades identificadas desde la perspectiva docente, se
propone el desarrollo de una máquina virtual educativa basada en una versión
simplificada de la arquitectura ARM32, acompañada de un assembler propio y
de documentación técnica diseñada como un manual de apoyo didáctico. Esta
propuesta busca ofrecer una herramienta que permita enseñar los principios
fundamentales de la arquitectura de un CPU sin depender de hardware físico ni
de una implementación completa y compleja de ARM.

Como complemento, se desarrollará un assembler propio encargado de
traducir programas escritos en lenguaje ensamblador a su representación
binaria. Esto permitirá a los estudiantes comprender la relación entre la sintaxis
del ensamblador, el formato de las instrucciones y su interpretación por parte
del procesador, reforzando así la conexión entre el software de bajo nivel y la
arquitectura subyacente.

Un componente central de la propuesta es la elaboración de documentación
técnica clara y estructurada, presentada en forma de manual. Dicho manual
describirá la arquitectura de la máquina virtual, el conjunto de instrucciones
soportadas, el formato de las mismas, el funcionamiento del assembler y
ejemplos prácticos de uso. De esta manera, el proyecto no solo proporcionará
una implementación funcional, sino también un recurso educativo reutilizable
por tutores y estudiantes.

Finalmente, el proyecto será implementado en el lenguaje de programación Go,
con el objetivo de garantizar su portabilidad y facilitar su ejecución en distintos
sistemas operativos. Esto permitirá que la herramienta pueda ser utilizada en
diversos entornos educativos sin requerir configuraciones complejas ni
dependencias de hardware específico.

== Cronograma

// #figure(image("img/gant.png", width: 95%),)

#text(font: "Comic Neue", size: 10pt)[
  #table(
    inset: 8pt,
    columns: (19%, auto, auto, auto, auto, auto, auto, auto),
    rows: (auto, 50pt, 50pt, 50pt, 50pt, 50pt, 50pt,),
    align: (left, left, left),
    table.header(
      table.cell(align: center+horizon, fill: header_bcolor,)[#text(style:"italic")[Actividades]],
      table.cell(align: center+horizon, fill: header_bcolor)[#text(style:"italic")[Enero 22-31]],
      table.cell(align: center+horizon, fill: header_bcolor)[#text(style:"italic")[Febrero 1-15]],
      table.cell(align: center+horizon, fill: header_bcolor)[#text(style:"italic")[Febrero 16-31]],
      table.cell(align: center+horizon, fill: header_bcolor)[#text(style:"italic")[Marzo 1-15]],
      table.cell(align: center+horizon, fill: header_bcolor)[#text(style:"italic")[Marzo 16-31]],
      table.cell(align: center+horizon, fill: header_bcolor)[#text(style:"italic")[Abril 1-15]],
      table.cell(align: center+horizon, fill: header_bcolor)[#text(style:"italic")[Abril 16-31]],
    ),
    table.cell(fill: header_bcolor,)[Analisis y diseño], table.cell(fill: red)[], [], [], [], [], [], [],
    table.cell(fill: header_bcolor)[Desarollo del assembler],[],table.cell(fill: red)[],table.cell(fill: red)[],table.cell(fill: red)[],[],[],[],
    table.cell(fill: header_bcolor)[Implementacion de la maquina virtual],[],[],table.cell(fill: red)[],table.cell(fill: red)[],table.cell(fill: red)[],[],[],
    table.cell(fill: header_bcolor)[Elaboracion de la documentacion],[],[],[],[],table.cell(fill: red)[],table.cell(fill:red)[],table.cell(fill: red)[],
    table.cell(fill: header_bcolor)[Elaboracionde integracion y pruebas],[],[],[],[],[],table.cell(fill: red)[],table.cell(fill: red)[],
    table.cell(fill: header_bcolor)[Refinamiento],[],[],[],[],[],[],table.cell(fill: red)[],
  )
]



=== Análisis y Diseño

En esta fase se realizará el análisis conceptual del proyecto, definiendo la
arquitectura general de la máquina virtual ARM32 simplificada. Se establecerá
el modelo de registros, el esquema de memoria, el conjunto inicial de
instrucciones y el formato básico de las mismas. Asimismo, se definirá la
sintaxis del lenguaje ensamblador y la estructura general de la documentación.

=== Implementación de la máquina virtual

Durante esta etapa se desarrollará la máquina virtual, implementando el ciclo
de ejecución (fetch–decode–execute), el manejo de registros de 32 bits, la
memoria lineal compartida y los flags de estado. El objetivo es obtener una
versión funcional que permite ejecutar instrucciones básicas y sirva como base
para futuras mejoras.

=== Desarrollo del assembler

En esta fase se implementará el assembler encargado de traducir programas
escritos en lenguaje ensamblador a palabras binarias de 32 bits compatibles
con la máquina virtual. Se buscará mantener coherencia entre el formato de las
instrucciones y su interpretación por la máquina virtual, permitiendo
modificaciones en ambos componentes cuando sea necesario.

=== Integración y pruebas

La máquina virtual y el assembler serán integrados y evaluados mediante
programas de prueba. Esta fase estará orientada a verificar el correcto
funcionamiento del sistema, identificar errores y validar que el comportamiento
de la máquina virtual sea coherente con el diseño propuesto.

=== Elaboración de la documentació

Se elaborará la documentación técnica en forma de manual, describiendo la
arquitectura del sistema, el conjunto de instrucciones, el funcionamiento del
assembler y ejemplos prácticos de uso. La documentación se desarrollará de
manera paralela al proyecto y se actualizará conforme se realicen cambios en
el diseño o la implementación.

=== Refinamiento

En la fase final se realizaron ajustes y mejoras al sistema en su conjunto, tanto
a nivel de implementación como de documentación. Esta etapa permitirá
optimizar decisiones de diseño, mejorar la claridad del manual y consolidar el
proyecto como una herramienta educativa coherente y funcional.

= Diseño y Arquitectura del Proyecto

== Arquitectura General del Sistema

La arquitectura general del proyecto se organiza como una cadena de etapas
claramente definidas, cuyo objetivo es facilitar la comprensión del flujo de
ejecución de un programa desde su escritura hasta su interpretación por la
máquina virtual.

El flujo del sistema comienza con la escritura de un programa en lenguaje
ensamblador utilizando un editor de texto de preferencia del usuario. Este
programa es posteriormente procesado por el ensamblador (assembler), el cual
analiza las instrucciones simbólicas y genera como resultado una secuencia de
palabras binarias de 32 bits, coherentes con el modelo de arquitectura definido
para el sistema.

Las palabras binarias generadas por el ensamblador son cargadas en la memoria
de la máquina virtual, donde representan tanto las instrucciones como los datos,
Una vez cargado el programa,
la máquina virtual inicia su ejecución aplicando el ciclo clásico de
fetch–decode–execute, en el cual se obtiene la instrucción desde memoria, se
decodifica su significado y se ejecuta la operación correspondiente sobre los
registros y la memoria del sistema.

Con el objetivo de reforzar el carácter educativo del proyecto, se incorpora
una interfaz gráfica de usuario (GUI) que permite visualizar de manera clara
el estado interno de la máquina virtual durante la ejecución del programa. A
través de esta interfaz es posible observar elementos como el contenido de los
registros, el estado de la memoria y la instrucción actualmente en ejecución,
facilitando el análisis del comportamiento del sistema paso a paso.

Esta arquitectura busca priorizar la claridad conceptual y la trazabilidad del
proceso de ejecución, permitiendo al usuario comprender la relación entre el
código ensamblador, su representación binaria y su efecto directo sobre el
estado interno de la máquina virtual.

== Arquitectura del Ensamblador 

El ensamblador constituye el primer componente funcional del sistema y es el
encargado de transformar el programa escrito en lenguaje ensamblador en una
representación binaria ejecutable por la máquina virtual. Su diseño se enfoca
en la claridad del proceso de traducción más que en la optimización extrema,
con fines principalmente educativos.

=== Syntaxis General

La sintaxis del assembler está basada en convenciones utilizadas en sistemas reales. Cada línea puede contener una instrucción, una etiqueta o un
comentario

==== Commentarios
```
// single line comments 
;  single line comments 
/* Multiline 
 * comments
 */
```

==== Etiquetas (Lables)

#lorem(30)

==== Valores inmediatios

#lorem(30)

=== Ensamblador de Dos Pasadas

El ensamblador implementado en este proyecto sigue un modelo de dos pasadas,
una estrategia clásica que permite resolver referencias simbólicas y generar
código binario correcto sin incrementar innecesariamente la complejidad del
sistema.

Durante la primera pasada, el ensamblador recorre el programa fuente con el
objetivo de analizar su estructura general y construir una tabla de símbolos.
En esta etapa se identifican etiquetas, se asignan direcciones de memoria a
cada instrucción y se valida la organización básica del código. No se genera
código binario definitivo en esta fase; en su lugar, se recopila la información
necesaria para la traducción posterior.

En la segunda pasada, el ensamblador vuelve a procesar el programa fuente,
utilizando la tabla de símbolos generada previamente para resolver referencias
a etiquetas y direcciones. En esta etapa se lleva a cabo la traducción completa
de las instrucciones a palabras binarias de 32 bits, ya con todos los operandos
resueltos y codificados según el formato de instrucción definido por la
arquitectura del sistema.

== Arquitectura de la maquina virtual

=== Arquitectura Von Neumann

La arquitectura Von Neumann se caracteriza por utilizar un único espacio de memoria
compartido tanto para las instrucciones como para los datos. En este modelo, el
procesador accede a la memoria a través de un solo bus, por el cual se transfieren
indistintamente instrucciones y datos.

Esta unificación simplifica el diseño del sistema, ya que no es necesario distinguir
entre distintos tipos de memoria ni entre diferentes rutas de acceso. No obstante,
introduce una limitación conocida como el cuello de botella de Von Neumann, debido
a que el procesador no puede acceder simultáneamente a instrucciones y datos,
lo que puede afectar el rendimiento.

A pesar de esta limitación, la arquitectura Von Neumann es ampliamente utilizada
en sistemas educativos y en modelos conceptuales por su simplicidad y claridad
estructural.

=== Arquitectura Harvard

La arquitectura Harvard se distingue por mantener memorias separadas para las
instrucciones y los datos, cada una con su propio bus de acceso. Esta separación
permite que el procesador obtenga una instrucción y acceda a datos al mismo tiempo,
mejorando el rendimiento del sistema.

Este modelo reduce el cuello de botella presente en la arquitectura Von Neumann y
permite optimizaciones como diferentes tamaños de palabra o tecnologías de memoria
para instrucciones y datos. Sin embargo, esta separación incrementa la complejidad
del diseño del hardware y del sistema en general.

=== Arquitectura de [NOMBRE PROYECTO]

Para este proyecto se decidio utilizar la arquitectura de X

== La Interfaz Gráfica de Usuario (GUI)

El objetivo
principal de esta interfaz es facilitar la comprensión del funcionamiento interno
de una arquitectura de CPU de forma visual e interactiva, especialmente con
fines educativos.

La interfaz mostrará de manera clara el estado de los registros del procesador,
permitiendo al estudiante observar cómo cambian sus valores conforme se
ejecutan las instrucciones. Esto es especialmente importante para entender
conceptos como la transferencia de datos, operaciones aritméticas y el uso de
registros generales.

Adicionalmente, la GUI incluirá una vista de la memoria, ubicada en un panel
lateral, donde se podrá observar el contenido de direcciones específicas. Esto
permitirá relacionar directamente las instrucciones del programa con su efecto
sobre la memoria, reforzando conceptos como direccionamiento, carga y
almacenamiento de datos.

Un elemento central de la interfaz será la ejecución paso a paso del programa.
Este modo permitirá avanzar instrucción por instrucción, actualizando en
tiempo real los registros, la memoria y el contador de programa. De esta
manera, el estudiante podrá analizar con detalle el efecto individual de cada
instrucción, lo cual resulta especialmente útil para el aprendizaje y la
depuración. Además del modo paso a paso, la interfaz permitirá la ejecución
continua del programa, manteniendo siempre sincronizada la visualización del
estado interno de la máquina virtual. En conjunto, la GUI no solo funcionará
como un medio de ejecución, sino como una herramienta didáctica que conecta
el código ensamblador con el comportamiento interno de la CPU.

#figure(
  image("img/mockup.png", width: 90%),
  //caption: [_Mock up del sistema grafico_],
  caption: [Prototipo conceptual de la interfaz gráfica del sistema],
  numbering: none
)

== Analisis Detallado

La arquitectura del proyecto se basa en una máquina virtual que modela una
versión simplificada de la arquitectura ARM32, diseñada específicamente con
fines educativos. El objetivo no es reproducir una implementación completa ni
compatible a nivel binario con ARM real, sino capturar los principios
fundamentales del diseño de un CPU.

=== Modelo de Central Processing Unit

La máquina virtual implementa un procesador con las siguientes
características

- Arquitectura de 32 bits, donde tanto las instrucciones como los datos se representan como palabras de 32 bits.
- Habrá 13 registros de propósito general

- Registros especiales, para *Stack Pointer*, *Link Register* y *Program Counter*

- Registro de banderas incluye indicadores fundamentales
  - *Zero (Z)*: se activa cuando el resultado de una operación es igual a cero.
  - *Negative (N)*: se activa cuando el bit más significativo del resultado es uno, indicando un valor negativo en representación con signo.
  - *Carry (C)*: se activa cuando ocurre un acarreo fuera del bit más significativo en una operación aritmética, o un préstamo en el caso de una resta.
  - *Overflow (V)*: se activa cuando el resultado de una operación con signo excede el rango representable, indicando un desbordamiento aritmético.

=== Memoria

El sistema cuenta con *CUANTA MEMORIA?*

*QUE TIPO DE ARQUITECTURA DECIDIMOS? HARDVARD/NUEMANN????*

=== Conjunto de Instrucciones (ISA)

El proyecto define un Instruction Set Architecture (ISA) reducido, inspirado en ARM32. Se implementará un subconjunto de instrucciones básicas, tales como:

- Instrucciones de movimiento de datos
- Operaciones aritméticas y lógicas
- Instrucciones de comparación
- Saltos condicionales y no condicionales
- Instrucciones de carga y almacenamiento

Cada instrucción será codificada explícitamente en formato binario, y la
máquina virtual se encargará de decodificar y reinterpretar esos bits como
instrucciones ARM simplificadas. Esto permite estudiar directamente el proceso de fetch–decode–execute.

//#let header_tcolor = rgb(245, 247, 250)
#text(font: "Comic Neue", size: 10pt)[
  #align(center)[
    #table(
      inset: 10pt,
      columns: (25%, 50%, 25%),
      align: (center, left, left),
      table.header(
        table.cell(align: center+horizon, fill: header_bcolor)[#text()[Instrucion]],
        table.cell(align: center+horizon, fill: header_bcolor)[#text()[Descripcion]],
        table.cell(align: center+horizon, fill: header_bcolor)[#text()[Ejemplo]],
      ),
      [mov], [Copia el valor de un registro a otro registro], [`mov Rd, Rs`],
      [mov], [Copia un numero a un registro], [`mov Rd, #42`],
      [add], [Suma valor de dos registro a registro resultante], [`add Rd, Rs1, Rs2`],
      [add], [Suma registro mas a valor, y lo guarda en registro], [`add Rd, Rs1, #42`],
    )
  ]
]
