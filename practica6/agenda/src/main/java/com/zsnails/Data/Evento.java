/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.zsnails.Data;

import java.util.Date;

/**
 *
 * @author oviquez
 */
public abstract class Evento implements ObjetoAgenda {
    private Date fecha;
    private String nombre;

    public Evento(Date fecha, String nombre) {
        this.fecha = fecha;
        this.nombre = nombre;
    }

    public Date getFecha() {
        return fecha;
    }

    public String getNombre() {
        return nombre;
    }

}
