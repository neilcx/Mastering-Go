class Solution {
public:
    vector<vector<int>> findDifference(vector<int>& nums1, vector<int>& nums2) {

        set<int> nums1_set, nums2_set;
        vector<vector<int>> answer(2,vector<int>());

        for(auto i : nums1){
            nums1_set.insert(i);
        }

        for(auto i : nums2){
            nums2_set.insert(i);
        }

        for(auto num: nums1_set){
            if (nums2_set.find(num) == nums2_set.end()){
                answer[0].push_back(num);
            }
        }

        for(auto num: nums2_set){
            if (nums1_set.find(num) == nums1_set.end()){
                answer[1].push_back(num);
            }
        }

        return answer;
    }
};