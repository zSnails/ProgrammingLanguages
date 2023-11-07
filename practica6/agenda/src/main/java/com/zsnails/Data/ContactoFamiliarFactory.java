package com.zsnails.Data;

/**
 * @author zSnails
 */
public class ContactoFamiliarFactory implements ContactoAbstractFactory {

    private String parentezco;
    private String nombre;
    private String telefono;

    public ContactoFamiliarFactory(String parentezco, String nombre, String telefono) {
        this.parentezco = parentezco;
        this.nombre = nombre;
        this.telefono = telefono;
    }

    @Override
    public Contacto createContacto() {
        return new ContactoFamiliar(parentezco, nombre, telefono);
    }

}
