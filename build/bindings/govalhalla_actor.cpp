

#include <stdexcept>
#include <sstream>
#include "baldr/rapidjson_utils.h"
#include "valhalla/proto/api.pb.h"
#include "tyr/actor.h"
#include "tyr/govalhalla_actor.h"
#include "valhalla/worker.h"  // For valhalla_exception_t


namespace govalhalla {

Actor::Actor(const std::string& config_json, bool auto_cleanup) 
        : actor_(std::make_unique<valhalla::tyr::actor_t>(parseConfig(config_json), auto_cleanup)) {}



std::pair<std::string, std::string> Actor::Act(valhalla::Api& api, const std::function<void()>* interrupt) {
    try {
        
        // auto response = actor_->act(api, interrupt);
        return std::make_pair(actor_->act(api, interrupt), std::string());
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }

std::pair<std::string, std::string> Actor::Route(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {

        // auto response = actor_->route(request, interrupt, api);
        return std::make_pair(std::string(actor_->route(request, interrupt, api)), std::string());
            
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }

std::pair<std::string, std::string> Actor::Locate(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        
        // auto response = actor_->locate(request, interrupt, api);
        return std::make_pair(actor_->locate(request, interrupt, api), std::string());
            
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }

std::pair<std::string, std::string> Actor::Matrix(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        
        // auto response = actor_->matrix(request, interrupt, api);
        return std::make_pair(actor_->matrix(request, interrupt, api), std::string());
            
            
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }

std::pair<std::string, std::string> Actor::OptimizedRoute(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        
        // auto response = actor_->optimized_route(request, interrupt, api);
        return std::make_pair(actor_->optimized_route(request, interrupt, api), std::string());
            
            
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }

std::pair<std::string, std::string> Actor::Isochrone(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        
        // auto response = actor_->isochrone(request, interrupt, api);
        return std::make_pair(actor_->isochrone(request, interrupt, api), std::string());
            
            
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }


std::pair<std::string, std::string> Actor::TraceRoute(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        
        // auto response = actor_->trace_route(request, interrupt, api);
        return std::make_pair(actor_->trace_route(request, interrupt, api), std::string());
            
            
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }


std::pair<std::string, std::string> Actor::TraceAttributes(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        
        // auto response = actor_->trace_attributes(request, interrupt, api);
        return std::make_pair(actor_->trace_attributes(request, interrupt, api), std::string());
            
            
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }


std::pair<std::string, std::string> Actor::Height(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        
        // auto response = actor_->height(request, interrupt, api);
        return std::make_pair(actor_->height(request, interrupt, api), std::string());
            
            
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }


std::pair<std::string, std::string> Actor::TransitAvailable(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        
        // auto response = actor_->transit_available(request, interrupt, api);
        return std::make_pair(actor_->transit_available(request, interrupt, api), std::string());
            
            
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }


std::pair<std::string, std::string> Actor::Expansion(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        
        // auto response = actor_->expansion(request, interrupt, api);
        return std::make_pair(actor_->expansion(request, interrupt, api), std::string());
            
            
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }


std::pair<std::string, std::string> Actor::Centroid(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        
        // auto response = actor_->centroid(request, interrupt, api);
        return std::make_pair(actor_->centroid(request, interrupt, api), std::string());
            
            
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }


std::pair<std::string, std::string> Actor::Status(const std::string& request, const std::function<void()>* interrupt, valhalla::Api* api) {
    try {
        
        // auto response = actor_->status(request, interrupt, api);
        return std::make_pair(actor_->status(request, interrupt, api), std::string());
            
            
        } catch (const valhalla::valhalla_exception_t& e) {
            return std::make_pair(std::string(), exception_to_json(e) );
        }catch (...) {
            return std::make_pair(std::string(), std::string("Unknown error"));
        }
        
        
    }

    
void Actor::Cleanup() {
    if (actor_) {
            actor_->cleanup();
        }
}

boost::property_tree::ptree Actor::parseConfig(const std::string& config_json) {

    // try {
        boost::property_tree::ptree pt;
        std::istringstream is(config_json);
        rapidjson::read_json(is, pt);  
        return pt;
        // }
    // catch (const valhalla::valhalla_exception_t& e) {
    //     throw std::runtime_error(std::string("Config parsing failed: ") );
    //     }
    }

// Helper function to convert valhalla_exception_t to a JSON string
std::string Actor::exception_to_json(const valhalla::valhalla_exception_t& ex) {
    std::ostringstream oss;

    // Start building the JSON object
    oss << "{\n";
    oss << " \"code\": " << ex.code << ",\n";
    oss << " \"message\": \"" << ex.what() << "\",\n";
    oss << " \"http_code\": " << ex.http_code << ",\n";
    oss << " \"http_message\": \"" << ex.http_message << "\",\n";
    oss << " \"osrm_error\": \"" << ex.osrm_error << "\"";

    // Add StatsD key if it's not empty
    if (!ex.statsd_key.empty()) {
        oss << ",\n \"statsd_key\": \"" << ex.statsd_key << "\"";
    }

    // Close the JSON object
    oss << "\n}";

    return oss.str();
}

} // namespace govalhalla
