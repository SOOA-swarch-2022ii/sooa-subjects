## sooa-subjects
Componente encargado de la gestión de las asignaturas y los cursos asociados a dichas asignaturas
http://localhost:PORT/sooa-sb-ms

### hechos:
- POST /new-subject
- POST /new-course
- GET  /subjects/all
- GET  /subjects/name={name}
- GET  /subjects/id={id}
- GET  /subjects/code={code}
- GET  /subjects/cam={campus}/fac={faculty}

- GET  /courses/all
- GET  /courses/id={id}
- GET  /courses/sb={subject}
- GET  /courses/sb={subject}/sm={semester}
- GET  /courses/profe={professor}
- GET  /courses/profe={professor}/sm={semester}
- GET  /courses/st={student}
- GET  /courses/st={student}/sm={semester}
- GET  /courses/sb={subject}/sm={semester}/sch/d={day}
- GET  /courses/location/house={house}
- GET  /courses/semester={sm}/d={day}/ti={ti}/tf={tf}

### por hacer:

