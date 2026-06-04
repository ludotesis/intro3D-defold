components {
  id: "Moneda"
  component: "/modelos/Moneda.model"
}
components {
  id: "Rotar"
  component: "/scripts/Rotar.script"
  properties {
    id: "velocidad_rotacion"
    value: "0.0, 0.0, 90.0"
    type: PROPERTY_TYPE_VECTOR3
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"default\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 0.5\n"
  "  data: 0.5\n"
  "  data: 0.25\n"
  "}\n"
  ""
}
