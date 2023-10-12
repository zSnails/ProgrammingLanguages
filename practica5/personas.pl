% vim:ft=prolog

persona("Brenda", "Martínez", [1, 0, 0, 1, 1, 1, 0, 0, 1, 0]).
persona("Carlos", "Hernández", [1, 0, 0, 1, 0, 1, 1, 1, 0, 1]).
persona("David", "López", [0, 1, 1, 0, 0, 1, 1, 0, 1, 0]).
persona("Elena", "Sánchez", [1, 0, 1, 0, 1, 0, 0, 1, 0, 1]).
persona("Fernando", "Gómez", [0, 1, 0, 1, 0, 0, 1, 1, 1, 0]).
persona("Gabriela", "Pérez", [1, 0, 1, 0, 0, 1, 1, 0, 1, 0]).
persona("Héctor", "Rodríguez", [0, 0, 1, 1, 1, 0, 1, 0, 1, 0]).
persona("Isabel", "Díaz", [1, 1, 0, 0, 1, 1, 0, 0, 1, 0]).
persona("Javier", "Gutiérrez", [0, 1, 0, 1, 0, 1, 1, 0, 0, 1]).
persona("Karla", "Núñez", [1, 1, 0, 1, 0, 1, 0, 0, 1, 0]).


similar([], [], 0).
similar([CAH|CAT], [CBH|CBT], Puntaje) :-
    similar(CAT, CBT, RPuntaje),
    ((CAH =:= CBH) -> Puntaje is 1 + RPuntaje; Puntaje is RPuntaje).

% NOTE: Uso, most_similar([<arreglo a buscar>], _, -1).
most_similar(Array, _, Score) :-
    persona(Name, LastName, PersonArray),
    similar(Array, PersonArray, CurrentScore),
    CurrentScore > Score,
    most_similar(Array, (Name, LastName, CurrentScore), CurrentScore),!.

most_similar(_, (Name, LastName, CurrentScore), _) :- 
    format("Most similar person is ~w ~w with a similarity score of ~w~n", [Name, LastName, CurrentScore]).
