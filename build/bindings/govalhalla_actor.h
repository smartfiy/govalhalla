
#ifndef GOVALHALLA_ACTOR_H
#define GOVALHALLA_ACTOR_H

#include "boost/property_tree/json_parser.hpp"
#include "valhalla/tyr/actor.h"
#include "valhalla/tyr/actor.h"
#include "valhalla/tyr/actor.h"
#include "valhalla/baldr/graphreader.h"
#include "valhalla/proto/api.pb.h"
#include "boost/property_tree/ptree.hpp"
#include "valhalla/baldr/rapidjson_utils.h"
#include "memory"


namespace govalhalla {

  class Actor {
  public:
    /**
     * Constructor
     * @param config         used to configure loki/thor/odin workers and their graphreaders
     * @param auto_cleanup   whether or not to auto clean workers after each action call
     */
    explicit Actor(const std::string& config_json, bool auto_cleanup = false);
     // Prevent copying due to unique_ptr
    Actor(const Actor&) = delete;
    Actor& operator=(const Actor&) = delete;
        
    // Move operations are automatically defaulted correctly
    Actor(Actor&&) = default;
    Actor& operator=(Actor&&) = default;
    
    /**
     * Manually clean the underlying workers
     */
    void Cleanup();

    /**
     * Perform the action specified in the options. This will both fill out the api object passed in and
     * return the serialized protobuf bytes if protobuf output was requested or json string if json
     * output was requested
     * @param api        object containing the request options and result after the call
     * @param interrupt  allows the underlying computation to be aborted via the functor throwing
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> Act(valhalla::Api& api, const std::function<void()>* interrupt = nullptr);

    /**
     * Perform the route action and return json or protobuf depending on which was requested. The
     * request may either be in the form of a json string provided by the request_str parameter or
     * contained in the api parameter as a deserialized protobuf object
     * @param request_str  json string if json input is being used empty otherwise
     * @param interrupt    allows the underlying computation to be aborted via the functor throwing
     * @param api          protobuffer object which can contain the input request via the options object
     *                     and will be filled out as the request is processed
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> Route(const std::string& request_str,
                      const std::function<void()>* interrupt = nullptr,
                      valhalla::Api* api = nullptr);

    /**
     * Perform the locate action and return json or protobuf depending on which was requested. The
     * request may either be in the form of a json string provided by the request_str parameter or
     * contained in the api parameter as a deserialized protobuf object
     * @param request_str  json string if json input is being used empty otherwise
     * @param interrupt    allows the underlying computation to be aborted via the functor throwing
     * @param api          protobuffer object which can contain the input request via the options object
     *                     and will be filled out as the request is processed
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> Locate(const std::string& request_str,
                      const std::function<void()>* interrupt = nullptr,
                      valhalla::Api* api = nullptr);

    /**
     * Perform the matrix action and return json or protobuf depending on which was requested. The
     * request may either be in the form of a json string provided by the request_str parameter or
     * contained in the api parameter as a deserialized protobuf object
     * @param request_str  json string if json input is being used empty otherwise
     * @param interrupt    allows the underlying computation to be aborted via the functor throwing
     * @param api          protobuffer object which can contain the input request via the options object
     *                     and will be filled out as the request is processed
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> Matrix(const std::string& request_str,
                      const std::function<void()>* interrupt = nullptr,
                      valhalla::Api* api = nullptr);

    /**
     * Perform the optimized_route action and return json or protobuf depending on which was requested.
     * The request may either be in the form of a json string provided by the request_str parameter or
     * contained in the api parameter as a deserialized protobuf object
     * @param request_str  json string if json input is being used empty otherwise
     * @param interrupt    allows the underlying computation to be aborted via the functor throwing
     * @param api          protobuffer object which can contain the input request via the options object
     *                     and will be filled out as the request is processed
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> OptimizedRoute(const std::string& request_str,
                                const std::function<void()>* interrupt = nullptr,
                                valhalla::Api* api = nullptr);

    /**
     * Perform the isochrone action and return json or protobuf depending on which was requested. The
     * request may either be in the form of a json string provided by the request_str parameter or
     * contained in the api parameter as a deserialized protobuf object
     * @param request_str  json string if json input is being used empty otherwise
     * @param interrupt    allows the underlying computation to be aborted via the functor throwing
     * @param api          protobuffer object which can contain the input request via the options object
     *                     and will be filled out as the request is processed
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> Isochrone(const std::string& request_str,
                          const std::function<void()>* interrupt = nullptr,
                          valhalla::Api* api = nullptr);

    /**
     * Perform the trace_route action and return json or protobuf depending on which was requested. The
     * request may either be in the form of a json string provided by the request_str parameter or
     * contained in the api parameter as a deserialized protobuf object
     * @param request_str  json string if json input is being used empty otherwise
     * @param interrupt    allows the underlying computation to be aborted via the functor throwing
     * @param api          protobuffer object which can contain the input request via the options object
     *                     and will be filled out as the request is processed
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> TraceRoute(const std::string& request_str,
                            const std::function<void()>* interrupt = nullptr,
                            valhalla::Api* api = nullptr);

    /**
     * Perform the trace_attributes action and return json or protobuf depending on which was requested.
     * The request may either be in the form of a json string provided by the request_str parameter or
     * contained in the api parameter as a deserialized protobuf object
     * @param request_str  json string if json input is being used empty otherwise
     * @param interrupt    allows the underlying computation to be aborted via the functor throwing
     * @param api          protobuffer object which can contain the input request via the options object
     *                     and will be filled out as the request is processed
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> TraceAttributes(const std::string& request_str,
                                const std::function<void()>* interrupt = nullptr,
                                valhalla::Api* api = nullptr);

    /**
     * Perform the height action and return json or protobuf depending on which was requested. The
     * request may either be in the form of a json string provided by the request_str parameter or
     * contained in the api parameter as a deserialized protobuf object
     * @param request_str  json string if json input is being used empty otherwise
     * @param interrupt    allows the underlying computation to be aborted via the functor throwing
     * @param api          protobuffer object which can contain the input request via the options object
     *                     and will be filled out as the request is processed
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> Height(const std::string& request_str,
                      const std::function<void()>* interrupt = nullptr,
                      valhalla::Api* api = nullptr);

    /**
     * Perform the transit_available action and return json or protobuf depending on which was
     * requested. The request may either be in the form of a json string provided by the request_str
     * parameter or contained in the api parameter as a deserialized protobuf object
     * @param request_str  json string if json input is being used empty otherwise
     * @param interrupt    allows the underlying computation to be aborted via the functor throwing
     * @param api          protobuffer object which can contain the input request via the options object
     *                     and will be filled out as the request is processed
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> TransitAvailable(const std::string& request_str,
                                  const std::function<void()>* interrupt = nullptr,
                                  valhalla::Api* api = nullptr);

    /**
     * Perform the expansion action and return json or protobuf depending on which was requested. The
     * request may either be in the form of a json string provided by the request_str parameter or
     * contained in the api parameter as a deserialized protobuf object
     * @param request_str  json string if json input is being used empty otherwise
     * @param interrupt    allows the underlying computation to be aborted via the functor throwing
     * @param api          protobuffer object which can contain the input request via the options object
     *                     and will be filled out as the request is processed
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> Expansion(const std::string& request_str,
                          const std::function<void()>* interrupt = nullptr,
                          valhalla::Api* api = nullptr);

    /**
     * Perform the centroid action and return json or protobuf depending on which was requested. The
     * request may either be in the form of a json string provided by the request_str parameter or
     * contained in the api parameter as a deserialized protobuf object
     * @param request_str  json string if json input is being used empty otherwise
     * @param interrupt    allows the underlying computation to be aborted via the functor throwing
     * @param api          protobuffer object which can contain the input request via the options object
     *                     and will be filled out as the request is processed
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> Centroid(const std::string& request_str,
                        const std::function<void()>* interrupt = nullptr,
                        valhalla::Api* api = nullptr);

    /**
     * Perform the status action and return json or protobuf depending on which was requested. The
     * request may either be in the form of a json string provided by the request_str parameter or
     * contained in the api parameter as a deserialized protobuf object
     * @param request_str  json string if json input is being used empty otherwise
     * @param interrupt    allows the underlying computation to be aborted via the functor throwing
     * @param api          protobuffer object which can contain the input request via the options object
     *                     and will be filled out as the request is processed
     * @return json or pbf bytes depending on what was specified in the options object
     */
    std::pair<std::string, std::string> Status(const std::string& request_str,
                      const std::function<void()>* interrupt = nullptr,
                      valhalla::Api* api = nullptr);

  private:
    static boost::property_tree::ptree parseConfig(const std::string& config_json);
    std::unique_ptr<valhalla::tyr::actor_t> actor_;


  };

} // namespace govalhalla

#endif // GOVALHALLA_ACTOR_H
