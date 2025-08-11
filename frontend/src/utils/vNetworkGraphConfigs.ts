import { defineConfigs } from "v-network-graph";
import {
  ForceLayout,
  type ForceNodeDatum,
  type ForceEdgeDatum,
} from "v-network-graph/lib/force-layout";

export function createGraphConfig(
  theme: "light" | "dark",
  isAddingEdge: boolean | number = false,
  resolution: boolean = false
) {
  const isDark = theme == "dark";

  return defineConfigs({
    view: {
      grid: {
        visible: true,
        interval: 10,
        thickIncrements: 5,
        line: {
          color: isDark ? "#444" : "#e0e0e0",
          width: 1,
          dasharray: 1,
        },
        thick: {
          color: isDark ? "#666" : "#ccc",
          width: 1,
          dasharray: 0,
        },
      },
      layoutHandler: new ForceLayout({
        positionFixedByDrag: false,
        positionFixedByClickWithAltKey: true,
        createSimulation: resolution
          ? (d3, nodes, edges) => {
              const forceLink = d3
                .forceLink<ForceNodeDatum, ForceEdgeDatum>(edges)
                .id((d: { id: any }) => d.id)
                .distance(100) // Increased distance for better spacing
                .strength(0.8); // Stronger link force for better structure

              return d3
                .forceSimulation(nodes)
                .force("edge", forceLink)
                .force("charge", d3.forceManyBody().strength(-300)) // Stronger repulsion for spacing
                .force("center", d3.forceCenter()) // Keep graph centered
                .force("collision", d3.forceCollide().radius(30)) // Prevent node overlap
                .alphaMin(0.001) // Fine-tune stopping condition
                .alphaDecay(0.0228) // Control simulation decay rate
                .velocityDecay(0.4); // Control velocity decay
            }
          : (d3, nodes, edges) => {
              const forceLink = d3
                .forceLink<ForceNodeDatum, ForceEdgeDatum>(edges)
                .id((d: { id: any }) => d.id);
              return d3
                .forceSimulation(nodes)
                .force("edge", forceLink.distance(100).strength(0.5))
                .force("charge", d3.forceManyBody().strength(-0.05))
                .alphaMin(0.001);
            },
      }),
      minZoomLevel: 0.5,
      maxZoomLevel: 10,
    },
    node: {
      selectable: isAddingEdge ? 2 : 1,
      normal: {
        color: isDark ? "#ff6699" : "#d13b6f",
        radius: 18,
      },
      hover: {
        color: isDark ? "#ff99cc" : "#f273a3",
        radius: 20,
      },
      label: {
        visible: true,
        direction: "south",
        directionAutoAdjustment: true,
        color: isDark ? "white" : "black",
      },
      focusring: {
        color: isAddingEdge ? "blue" : "yellow",
      },
      zOrder: {
        enabled: true,
        zIndex: (n) => n.zIndex,
        bringToFrontOnHover: true,
        bringToFrontOnSelected: true,
      },
    },
    edge: {
      selectable: false,
      gap: 50,
      normal: {
        color: isDark ? "#ffb3cc" : "#f18ca3",
        width: 4,
      },
      hover: {
        color: isDark ? "#ff99bb" : "#ff6f96",
        width: 8,
      },
      selected: {
        width: 6,
      },
      label: {
        color: isDark ? "white" : "black",
      },
      type: "curve",
      zOrder: {
        enabled: true,
        zIndex: (n) => n.zIndex,
        bringToFrontOnHover: true,
        bringToFrontOnSelected: true,
      },
    },
    path: resolution
      ? {
          visible: true,
          path: {
            width: 10,
            dasharray: "10 16",
            animate: true,
            animationSpeed: 40,
            color: "red",
          },
        }
      : {},
  });
}
