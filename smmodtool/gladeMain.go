package main

var gladeMain = `<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated with glade 3.22.1 -->
<interface>
  <requires lib="gtk+" version="3.20"/>
  <object class="GtkWindow" id="window">
    <property name="width_request">900</property>
    <property name="height_request">600</property>
    <property name="can_focus">False</property>
    <property name="border_width">3</property>
    <child type="titlebar">
      <object class="GtkHeaderBar">
        <property name="visible">True</property>
        <property name="can_focus">False</property>
        <property name="show_close_button">True</property>
        <child>
          <object class="GtkButtonBox">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="homogeneous">True</property>
            <property name="layout_style">start</property>
            <child>
              <object class="GtkButton" id="buttonSave">
                <property name="label" translatable="yes">Save</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="buttonLoad">
                <property name="label" translatable="yes">Open</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
            <child>
              <object class="GtkButton" id="buttonCompile">
                <property name="label" translatable="yes">Compile</property>
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="receives_default">True</property>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">2</property>
              </packing>
            </child>
          </object>
        </child>
      </object>
    </child>
    <child>
      <object class="GtkBox">
        <property name="visible">True</property>
        <property name="can_focus">False</property>
        <child>
          <object class="GtkBox">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="orientation">vertical</property>
            <child>
              <object class="GtkScrolledWindow">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="hscrollbar_policy">never</property>
                <property name="vscrollbar_policy">always</property>
                <property name="shadow_type">in</property>
                <child>
                  <object class="GtkViewport">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <child>
                      <object class="GtkListBox" id="listParts">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <child>
                          <object class="GtkListBoxRow">
                            <property name="visible">True</property>
                            <property name="can_focus">True</property>
                            <child>
                              <object class="GtkLabel">
                                <property name="visible">True</property>
                                <property name="can_focus">False</property>
                                <property name="events"/>
                                <property name="label" translatable="yes">Part name aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</property>
                                <property name="ellipsize">end</property>
                                <property name="single_line_mode">True</property>
                                <attributes>
                                  <attribute name="weight" value="bold"/>
                                </attributes>
                              </object>
                            </child>
                          </object>
                        </child>
                      </object>
                    </child>
                  </object>
                </child>
              </object>
              <packing>
                <property name="expand">True</property>
                <property name="fill">True</property>
                <property name="position">0</property>
              </packing>
            </child>
            <child>
              <object class="GtkBox">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <child>
                  <object class="GtkButton" id="buttonAddPart">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <property name="margin_left">2</property>
                    <property name="margin_right">2</property>
                    <property name="margin_top">2</property>
                    <property name="margin_bottom">2</property>
                    <property name="always_show_image">True</property>
                    <child>
                      <object class="GtkImage">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="stock">gtk-add</property>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">False</property>
                    <property name="fill">True</property>
                    <property name="position">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkSearchEntry" id="searchPart">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="margin_left">2</property>
                    <property name="margin_right">2</property>
                    <property name="margin_top">2</property>
                    <property name="margin_bottom">2</property>
                    <property name="primary_icon_name">edit-find-symbolic</property>
                    <property name="primary_icon_activatable">False</property>
                    <property name="primary_icon_sensitive">False</property>
                  </object>
                  <packing>
                    <property name="expand">False</property>
                    <property name="fill">True</property>
                    <property name="position">1</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkButton" id="buttonDeletePart">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="receives_default">True</property>
                    <property name="margin_left">2</property>
                    <property name="margin_right">2</property>
                    <property name="margin_top">2</property>
                    <property name="margin_bottom">2</property>
                    <property name="always_show_image">True</property>
                    <child>
                      <object class="GtkImage">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="stock">gtk-delete</property>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">False</property>
                    <property name="fill">True</property>
                    <property name="position">2</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="expand">False</property>
                <property name="fill">True</property>
                <property name="position">1</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">False</property>
            <property name="fill">True</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkNotebook">
            <property name="name">notebook</property>
            <property name="visible">True</property>
            <property name="can_focus">True</property>
            <child>
              <object class="GtkBox">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_left">5</property>
                <property name="margin_right">5</property>
                <property name="margin_top">5</property>
                <property name="margin_bottom">5</property>
                <property name="orientation">vertical</property>
                <child>
                  <object class="GtkBox">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="margin_bottom">5</property>
                    <child>
                      <object class="GtkLabel">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="margin_right">5</property>
                        <property name="label" translatable="yes">Name</property>
                      </object>
                      <packing>
                        <property name="expand">False</property>
                        <property name="fill">True</property>
                        <property name="position">0</property>
                      </packing>
                    </child>
                    <child>
                      <object class="GtkEntry">
                        <property name="visible">True</property>
                        <property name="can_focus">True</property>
                        <property name="margin_right">5</property>
                      </object>
                      <packing>
                        <property name="expand">True</property>
                        <property name="fill">True</property>
                        <property name="position">1</property>
                      </packing>
                    </child>
                    <child>
                      <object class="GtkLabel">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="margin_right">5</property>
                        <property name="label" translatable="yes">Language</property>
                      </object>
                      <packing>
                        <property name="expand">False</property>
                        <property name="fill">True</property>
                        <property name="position">2</property>
                      </packing>
                    </child>
                    <child>
                      <object class="GtkComboBoxText" id="descriptionLanguage">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="active">1</property>
                        <items>
                          <item id="capitalist" translatable="yes">English</item>
                          <item id="hitler" translatable="yes">German</item>
                          <item id="vodka" translatable="yes">Russian</item>
                          <item id="white" translatable="yes">French</item>
                          <item id="neigh" translatable="yes">Equestrian</item>
                        </items>
                      </object>
                      <packing>
                        <property name="expand">False</property>
                        <property name="fill">True</property>
                        <property name="position">3</property>
                      </packing>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">False</property>
                    <property name="fill">True</property>
                    <property name="position">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkScrolledWindow">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="shadow_type">in</property>
                    <child>
                      <object class="GtkTextView">
                        <property name="visible">True</property>
                        <property name="can_focus">True</property>
                        <property name="margin_left">4</property>
                        <property name="margin_right">4</property>
                        <property name="margin_top">4</property>
                        <property name="margin_bottom">4</property>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">True</property>
                    <property name="fill">True</property>
                    <property name="position">1</property>
                  </packing>
                </child>
              </object>
            </child>
            <child type="tab">
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="label" translatable="yes">Description</property>
              </object>
              <packing>
                <property name="tab_fill">False</property>
              </packing>
            </child>
            <child>
              <object class="GtkBox">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="margin_left">5</property>
                <property name="margin_right">5</property>
                <property name="margin_top">5</property>
                <property name="margin_bottom">5</property>
                <property name="orientation">vertical</property>
                <child>
                  <object class="GtkBox">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="margin_bottom">5</property>
                    <child>
                      <object class="GtkLabel">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="label" translatable="yes">UUID</property>
                      </object>
                      <packing>
                        <property name="expand">False</property>
                        <property name="fill">True</property>
                        <property name="position">0</property>
                      </packing>
                    </child>
                    <child>
                      <object class="GtkEntry">
                        <property name="visible">True</property>
                        <property name="can_focus">True</property>
                      </object>
                      <packing>
                        <property name="expand">True</property>
                        <property name="fill">True</property>
                        <property name="position">1</property>
                      </packing>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">False</property>
                    <property name="fill">True</property>
                    <property name="position">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkScrolledWindow">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="shadow_type">in</property>
                    <child>
                      <object class="GtkTextView">
                        <property name="visible">True</property>
                        <property name="can_focus">True</property>
                        <property name="margin_top">4</property>
                        <property name="margin_bottom">4</property>
                      </object>
                    </child>
                  </object>
                  <packing>
                    <property name="expand">True</property>
                    <property name="fill">True</property>
                    <property name="position">1</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="position">1</property>
              </packing>
            </child>
            <child type="tab">
              <object class="GtkLabel">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="label" translatable="yes">Part data</property>
              </object>
              <packing>
                <property name="position">1</property>
                <property name="tab_fill">False</property>
              </packing>
            </child>
          </object>
          <packing>
            <property name="expand">True</property>
            <property name="fill">True</property>
            <property name="position">1</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
</interface>`