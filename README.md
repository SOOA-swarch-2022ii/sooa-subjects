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
- GET  /courses/all
- GET  /courses/id={id}
- GET  /courses/sb={subject_code}
- GET  /courses/st={student}
- GET  /courses/profe={professor}



### por hacer:
- PUT /courses/id={id}/st={student}/add
- PUT /courses/id={id}/st={student}/remove

- GET  /courses/st={student}/sm={semester}
- GET  /subjects/cam={campus}
- GET  /subjects/cam={campus}/fac={faculty}
- GET  /subjects/cam={campus}/fac={faculty}/bau={bau}
- GET  /courses/sb={subject}/sm={semester}
- GET  /courses/d={day}/ti={ti}/tf={tf}
- GET  /courses/profe={professor}/sem={semester}
- GET  /courses/{location}