package com.zsnails.Data;

import java.sql.Date;

/**
 * @author zSnails
 */
public class EventoReunionFactory implements EventoAbstractFactory {

    private int cantidadAsistentes;
    private Date fecha;
    private String nombre;

    public EventoReunionFactory(int cantidadAsistentes, Date fecha, String nombre) {
        this.cantidadAsistentes = cantidadAsistentes;
        this.fecha = fecha;
        this.nombre = nombre;
    }

    @Override
    public Evento createEvento() {
        return new EventoReunion(cantidadAsistentes, fecha, nombre);
    }

}
