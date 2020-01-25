package main

var gladeMain = `<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated with glade 3.22.1 -->
<interface>
  <requires lib="gtk+" version="3.20"/>
  <object class="GtkListStore" id="listStoreParts">
    <columns>
      <!-- column-name Parts -->
      <column type="gchararray"/>
    </columns>
    <data>
      <row>
        <col id="0" translatable="yes">Hello world</col>
      </row>
    </data>
  </object>
  <object class="GtkWindow" id="windowMain">
    <property name="width_request">900</property>
    <property name="height_request">600</property>
    <property name="can_focus">False</property>
    <property name="border_width">3</property>
    <child type="titlebar">
      <object class="GtkHeaderBar">
        <property name="visible">True</property>
        <property name="can_focus">False</property>
        <property name="title" translatable="yes">S410's SM Mod Tool</property>
        <property name="subtitle" translatable="yes">https://github.com/Sierra410/sm-tool</property>
        <property name="spacing">10</property>
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
        <child>
          <object class="GtkStackSwitcher" id="stackSwitcherLog">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="stack">stackLog</property>
          </object>
          <packing>
            <property name="pack_type">end</property>
            <property name="position">1</property>
          </packing>
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
                <property name="shadow_type">in</property>
                <child>
                  <object class="GtkViewport">
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <child>
                      <object class="GtkListBox" id="listParts">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
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
                  <object class="GtkSearchEntry" id="searchEntryPart">
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
          <object class="GtkStack" id="stackLog">
            <property name="visible">True</property>
            <property name="can_focus">False</property>
            <property name="transition_duration">500</property>
            <property name="transition_type">slide-up</property>
            <child>
              <object class="GtkStack" id="stackPartData">
                <property name="visible">True</property>
                <property name="can_focus">False</property>
                <property name="transition_duration">500</property>
                <property name="transition_type">slide-left</property>
                <child>
                  <object class="GtkLabel">
                    <property name="name">0</property>
                    <property name="visible">True</property>
                    <property name="can_focus">False</property>
                    <property name="label" translatable="yes">Select a part</property>
                  </object>
                  <packing>
                    <property name="name">0</property>
                  </packing>
                </child>
                <child>
                  <object class="GtkNotebook" id="notebookPartData">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <child>
                      <object class="GtkBox">
                        <property name="visible">True</property>
                        <property name="can_focus">False</property>
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkBox">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="margin_left">5</property>
                            <property name="margin_right">5</property>
                            <property name="margin_top">5</property>
                            <property name="margin_bottom">5</property>
                            <child>
                              <object class="GtkLabel">
                                <property name="visible">True</property>
                                <property name="can_focus">False</property>
                                <property name="margin_left">5</property>
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
                              <object class="GtkEntry" id="entryName">
                                <property name="visible">True</property>
                                <property name="can_focus">True</property>
                                <property name="caps_lock_warning">False</property>
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
                                <property name="margin_left">5</property>
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
                              <object class="GtkComboBoxText" id="comboBoxLanguage">
                                <property name="visible">True</property>
                                <property name="can_focus">False</property>
                                <property name="active">2</property>
                                <items>
                                  <item id="Brazilian" translatable="yes">Brazilian</item>
                                  <item id="Chinese" translatable="yes">Chinese</item>
                                  <item id="English" translatable="yes">English</item>
                                  <item id="French" translatable="yes">French</item>
                                  <item id="German" translatable="yes">German</item>
                                  <item id="Italian" translatable="yes">Italian</item>
                                  <item id="Japanese" translatable="yes">Japanese</item>
                                  <item id="Korean" translatable="yes">Korean</item>
                                  <item id="Polish" translatable="yes">Polish</item>
                                  <item id="Russian" translatable="yes">Russian</item>
                                  <item id="Spanish" translatable="yes">Spanish</item>
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
                            <property name="margin_left">5</property>
                            <property name="margin_right">5</property>
                            <property name="margin_top">5</property>
                            <property name="margin_bottom">5</property>
                            <property name="shadow_type">in</property>
                            <child>
                              <object class="GtkTextView" id="textViewDescription">
                                <property name="visible">True</property>
                                <property name="can_focus">True</property>
                                <property name="margin_left">5</property>
                                <property name="margin_right">5</property>
                                <property name="margin_top">5</property>
                                <property name="margin_bottom">5</property>
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
                        <property name="orientation">vertical</property>
                        <child>
                          <object class="GtkBox">
                            <property name="visible">True</property>
                            <property name="can_focus">False</property>
                            <property name="margin_left">5</property>
                            <property name="margin_top">5</property>
                            <property name="margin_bottom">5</property>
                            <child>
                              <object class="GtkLabel">
                                <property name="width_request">-1</property>
                                <property name="visible">True</property>
                                <property name="can_focus">False</property>
                                <property name="margin_left">5</property>
                                <property name="margin_right">5</property>
                                <property name="label" translatable="yes">UUID</property>
                              </object>
                              <packing>
                                <property name="expand">False</property>
                                <property name="fill">True</property>
                                <property name="position">0</property>
                              </packing>
                            </child>
                            <child>
                              <object class="GtkEntry" id="entryUuid">
                                <property name="visible">True</property>
                                <property name="can_focus">True</property>
                                <property name="caps_lock_warning">False</property>
                                <property name="progress_pulse_step">0.10000000022351742</property>
                              </object>
                              <packing>
                                <property name="expand">True</property>
                                <property name="fill">True</property>
                                <property name="position">1</property>
                              </packing>
                            </child>
                            <child>
                              <object class="GtkButton" id="buttonRandomUuid">
                                <property name="label" translatable="yes">Random</property>
                                <property name="visible">True</property>
                                <property name="can_focus">True</property>
                                <property name="receives_default">True</property>
                                <property name="margin_left">5</property>
                                <property name="margin_right">5</property>
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
                            <property name="position">0</property>
                          </packing>
                        </child>
                        <child>
                          <object class="GtkScrolledWindow">
                            <property name="visible">True</property>
                            <property name="can_focus">True</property>
                            <property name="margin_left">5</property>
                            <property name="margin_right">5</property>
                            <property name="margin_top">5</property>
                            <property name="margin_bottom">5</property>
                            <property name="shadow_type">in</property>
                            <child>
                              <object class="GtkTextView" id="textViewPartData">
                                <property name="visible">True</property>
                                <property name="can_focus">True</property>
                                <property name="margin_left">5</property>
                                <property name="margin_right">5</property>
                                <property name="margin_top">5</property>
                                <property name="margin_bottom">5</property>
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
                    <property name="name">1</property>
                    <property name="position">1</property>
                  </packing>
                </child>
              </object>
              <packing>
                <property name="name">0</property>
                <property name="title" translatable="yes">Part Data</property>
              </packing>
            </child>
            <child>
              <object class="GtkScrolledWindow">
                <property name="visible">True</property>
                <property name="can_focus">True</property>
                <property name="hscrollbar_policy">never</property>
                <property name="shadow_type">in</property>
                <child>
                  <object class="GtkTextView" id="textViewLog">
                    <property name="visible">True</property>
                    <property name="can_focus">True</property>
                    <property name="margin_left">10</property>
                    <property name="margin_right">10</property>
                    <property name="margin_top">10</property>
                    <property name="margin_bottom">10</property>
                    <property name="editable">False</property>
                    <property name="wrap_mode">char</property>
                    <property name="accepts_tab">False</property>
                  </object>
                </child>
              </object>
              <packing>
                <property name="name">1</property>
                <property name="title" translatable="yes">Log</property>
                <property name="position">1</property>
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
