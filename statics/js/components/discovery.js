/* jshint multistr: true */

var DiscoveryComponent = {

  name: 'discovery',

  mixins: [apiMixin],

  template: '\
    <div class="discovery">\
      <div class="col-sm-8 fill content">\
        <div id="sequence"></div>\
        <div class="discovery-d3"></div>\
      </div>\
      <div class="col-sm-4 fill sidebar">\
        <div class="left-cont">\
          <div class="left-panel">\
            <label for="mode">Type</label>\
            <select id="type" v-model="type" class="form-control input-sm">\
              <option value="bytes">Bytes</option>\
              <option value="packets">Packets</option>\
            </select>\
            <label for="mode">Mode</label>\
            <select id="mode" v-model="mode" class="form-control input-sm">\
              <option value="size">Size</option>\
              <option value="count">Count</option>\
            </select>\
          </div>\
          <div class="left-panel" v-if="protocolData">\
            <div class="title-left-panel">Protocol Data</div>\
            <object-detail :object="protocolData"></object-detail>\
          </div>\
        </div>\
      </div>\
  ',

  data: function() {
    return {
      protocolData: null,
      type: "bytes",
      mode: "count",
    };
  },

  mounted: function() {
    this.layout = new DiscoveryLayout(this, ".discovery-d3");
    this.layout.DrawChart(this.type);
  },

  watch: {

    type: function() {
      this.layout.DrawChart(this.type);
    },

    mode: function() {
      this.layout.ChangeMode(this.mode);
    }

  }

};

var DiscoveryLayout = function(vm, selector) {
  this.vm = vm;
  this.width = 680;
  this.height = 600;
  this.radius = (Math.min(this.width, this.height) / 2) - 50;
  this.color = d3.scaleOrdinal(d3.schemeCategory20);

  this.svg = d3.select(selector).append("svg")
    .attr("width", this.width)
    .attr("height", this.height)
    .append("g")
    .attr("id", "container")
    .attr("transform", "translate(" + this.width / 2 + "," + this.height * 0.52 + ")");

  this.partition = d3.partition()
    .size([2 * Math.PI, this.radius * this.radius]);

  var x = d3.scaleLinear()
      .range([0, 2 * Math.PI]);

  var y = d3.scaleSqrt()
      .range([0, this.radius]);

  this.arc = d3.arc()
      .startAngle(function(d) { return Math.max(0, Math.min(2 * Math.PI, x(d.x0))); })
      .endAngle(function(d) { return Math.max(0, Math.min(2 * Math.PI, x(d.x1))); })
      .innerRadius(function(d) { return Math.max(0, y(d.y0)); })
      .outerRadius(function(d) { return Math.max(0, y(d.y1)); });

  this.color = d3.scaleOrdinal(d3.schemeCategory20);

  this.partition = d3.partition();

  this.sum = function(d) { return 1; };

  // Breadcrumb dimensions: width, height, spacing, width of tip/tail.
  this.b = {
    w: 75, h: 30, s: 3, t: 10
  };
  this.initializeBreadcrumbTrail();
};

DiscoveryLayout.prototype.ChangeMode = function(mode) {
  var self = this;
  self.sum = mode === "count" ? function() { return 1; } : function(d) { return d.size; };

  // Interpolate the arcs in data space.
  function arcTween(a) {
    var i = d3.interpolate({x: a.x0, dx: a.dx0}, a);
    return function(t) {
      var b = i(t);
      a.x0 = b.x;
      a.dx0 = b.dx;
      return self.arc(b);
    };
  }

  var root = d3.hierarchy(self.root);
  root.sum(self.sum);
  this.path
    .data(self.partition(root).descendants())
    .transition()
    .duration(1500)
    .attrTween("d", arcTween);
  this.totalSize = this.path.datum().value;
};

DiscoveryLayout.prototype.DrawChart = function(type) {
  var self = this;
  var gremlinQuery = "g.Flows().Has('Link.Protocol', 'ETHERNET')";

  this.totalSize = 0;
  this.svg.selectAll("*").remove();
  this.vm.$topologyQuery(gremlinQuery)
    .then(function(data) {
      if (data === null)
        return [];

      var pathMap = {};
      data.forEach(function(flow, i) {
        var linkMetric = flow.Metric;
        var layersPath = flow.LayersPath;
        var metric = pathMap[layersPath];
        if (metric === undefined) {
          metric = { Bytes: 0, Packets: 0 };
        }
        metric.Bytes += linkMetric.ABBytes;
        metric.Bytes += linkMetric.BABytes;
        metric.Packets += linkMetric.ABPackets;
        metric.Packets += linkMetric.BAPackets;
        pathMap[layersPath] = metric;
      });

      var root = {name: "root", children: []};
      for (var path in pathMap) {
        var stats = pathMap[path];
        var node = root;
        var layers = path.split("/");
        for (var i in layers) {
          var l = undefined;
          var layer = layers[i];
          for (var c in node.children) {
            if (node.children[c].name == layer) {
              l = node.children[c];
              break;
            }
          }
          if (l === undefined) {
            l = {"name": layer, children: [], size: 0};
            node.children.push(l);
          }
          if (i == layers.length - 1) {
            l.size = type == "bytes" ? stats.Bytes : stats.Packets;
          }
          node = l;
        }
      }

      self.root = root
      root = d3.hierarchy(root);
      root.sum(self.sum);
      self.path = self.svg.selectAll("path")
          .data(self.partition(root).descendants())
        .enter().append("path")
          .attr("display", function(d) { return d.depth ? null : "none"; }) // hide inner ring
          .attr("d", self.arc)
          .style("stroke", "#fff")
          .style("fill", function(d) { return self.color((d.children ? d : d.parent).data.name); })
          .style("fill-rule", "evenodd")
          .on("mouseover", mouseover)
          .each(stash);
      self.totalSize = self.path.datum().value;
    });

  // On mouseover function
  function mouseover(d) {
    var percentage = (100 * d.value / self.totalSize).toPrecision(3) + " %";
    self.vm.protocolData = {
      "Name": d.data.name,
      "Percentage": percentage,
      "Size": d.data.size,
      "Value": d.value,
      "Depth": d.depth
    };
    var sequenceArray = getAncestors(d);
    updateBreadcrumbs(sequenceArray, percentage);
  }

  // On mouseleave function
  function mouseleave(d) {
    d3.select("#trail")
      .style("visibility", "hidden");
    self.vm.protocolData = null;
  }

  // Given a node in a partition layout, return an array of all of its ancestor
  // nodes, highest first, but excluding the root.
  function getAncestors(node) {
    var path = [];
    var current = node;
    while (current.parent) {
      path.unshift(current);
      current = current.parent;
    }
    return path;
  }

  // Generate a string that describes the points of a breadcrumb polygon.
  function breadcrumbPoints(d, i) {
    var points = [];
    points.push("0,0");
    points.push(self.b.w + ",0");
    points.push(self.b.w + self.b.t + "," + (self.b.h / 2));
    points.push(self.b.w + "," + self.b.h);
    points.push("0," + self.b.h);
    if (i > 0) { // Leftmost breadcrumb; don't include 6th vertex.
      points.push(self.b.t + "," + (self.b.h / 2));
    }
    return points.join(" ");
  }

  //Update the breadcrumb trail to show the current sequence and percentage.
  function updateBreadcrumbs(nodeArray, percentageString) {

    // Data join; key function combines name and depth (= position in sequence).
    var g = d3.select("#trail")
      .selectAll("g")
      .data(nodeArray, function(d) { return d.data.name + d.depth; });

    // Add breadcrumb and label for entering nodes.
    var entering = g.enter().append("svg:g");

    entering.append("svg:polygon")
      .attr("points", breadcrumbPoints)
      .style("fill", function(d) { return self.color((d.children ? d : d.parent).name); });

    entering.append("svg:text")
      .attr("x", (self.b.w + self.b.t) / 2)
      .attr("y", self.b.h / 2)
      .attr("dy", "0.35em")
      .attr("text-anchor", "middle")
      .text(function(d) { return d.data.name; });

    // Merge enter and update selections; set position for all nodes.
    entering.merge(g).attr("transform", function(d, i) {
      return "translate(" + i * (self.b.w + self.b.s) + ", 0)";
    });

    // Remove exiting nodes.
    g.exit().remove();

    // Now move and update the percentage at the end.
    d3.select("#trail").select("#endlabel")
      .attr("x", (nodeArray.length + 0.5) * (self.b.w + self.b.s))
      .attr("y", self.b.h / 2)
      .attr("dy", "0.35em")
      .attr("text-anchor", "middle")
      .text(percentageString);

    // Make the breadcrumb trail visible, if it's hidden.
    d3.select("#trail")
      .style("visibility", "");
  }

  // Stash the old values for transition.
  function stash(d) {
    d.x0 = d.x;
    d.dx0 = d.dx;
  }

  d3.select(self.frameElement).style("height", this.height + "px");
};

DiscoveryLayout.prototype.initializeBreadcrumbTrail = function() {
  // Add the svg area.
  var trail = d3.select("#sequence").append("svg:svg")
    .attr("width", this.width)
    .attr("height", 50)
    .attr("id", "trail");
  // Add the label at the end, for the percentage.
  trail.append("svg:text")
    .attr("id", "endlabel")
    .style("fill", "#fff");
};
