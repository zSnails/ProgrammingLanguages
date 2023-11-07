package com.zsnails.agenda;

import com.zsnails.Data.Agenda;
import com.zsnails.GUI.AgendaMain;

/**
 * Hello world!
 *
 */
public class App {
    // NOTE: La diferencia entre un eager singleton y un lazy singleton es que el
    // eager singleton se instancia apenas comienza el programa, un lazy singleton
    // se instancia cuando se necesita usar la primera vez, para este programa es
    // mejor un eager singleton, ya que, al menos en la
    // implementación que yo le di, se necesita tener una sola instancia de la
    // agenda en todo el programa, por lo que es buena idea inicializarlo justo
    // cuando inicia el programa.
    public static void main(String[] args) {
        Agenda miAgenda = Agenda.getInstance();
        AgendaMain guiMain = new AgendaMain(miAgenda);
        guiMain.setVisible(true);
    }
}
