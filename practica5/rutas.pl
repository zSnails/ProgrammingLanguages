% vim:ft=prolog
conectado(i,a,1).
conectado(i,b,2).
conectado(a,b,3).
conectado(a,c,4).
conectado(b,f,5).
conectado(b,c,6).
conectado(c,f,7).

peso(X, Y, P) :-
    conectado(X, Y, P), !.

conectados(X,Y) :- conectado(X,Y, _).
conectados(X,Y) :- conectado(Y,X, _).

shortest(Ini, Fin, Ruta) :-
    findall(R, ruta(Ini, Fin, R), Bag),
    sort(Bag, [(_, Ruta)|_]).


sum_list([], 0).
sum_list([H|T], Sum) :-
   sum_list(T, Rest),
   Sum is H + Rest.

get_total_weight([], 0).
get_total_weight([(_, W)|T], Weight) :- 
    get_total_weight(T, Rest),
    Weight is W + Rest.

ruta(Ini,Fin,(W, Ruta)) :- 
    ruta2(Ini,Fin,[],Ruta),
    get_total_weight(Ruta, W).

ruta2(Fin,Fin,_,[(Fin, 0)]).
ruta2(Ini,Fin,Visitados,[(Ini, P)|Resto]):-
    conectados(Ini,Alguien),
    peso(Ini, Alguien, P),
    not(member(Alguien,Visitados)),
    ruta2(Alguien,Fin,[Ini|Visitados],Resto).
