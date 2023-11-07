/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.zsnails.Data;

import java.util.LinkedList;

/**
 *
 * @author oviquez
 */
public class Agenda {
    private LinkedList<Object> listaObjetos;
    private static Agenda instance = null;

    private LinkedList<String> contactos;
    private LinkedList<String> eventos;

    public static Agenda getInstance() {
        if (instance == null) {
            instance = new Agenda();
        }
        return instance;
    }

    private Agenda() {
        this.listaObjetos = new LinkedList<Object>();
        this.contactos = new LinkedList<String>();
        this.eventos = new LinkedList<String>();
    }

    public LinkedList<String> getContactos() {
        return contactos;
    }

    public LinkedList<String> getEventos() {
        return eventos;
    }

    public LinkedList<Object> getListaObjetos() {
        return listaObjetos;
    }

}
